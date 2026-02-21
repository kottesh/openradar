package server

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golang.org/x/time/rate"
	"gorm.io/gorm"

	"openradar/app"

	"openradar/internal/config"

	"github.com/gorilla/websocket"
)

type ipLimiter struct {
	mu       sync.Mutex
	limiters map[string]*visitorLimiter
}

type visitorLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type Hub struct {
	clients   map[*websocket.Conn]bool
	mu        sync.Mutex
	Broadcast chan []byte
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { // TODO
		return true
	},
}

func newHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte, 256),
	}
}

func (h *Hub) run() {
	for msg := range h.Broadcast {
		h.mu.Lock()
		for conn := range h.clients {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				conn.Close()
				delete(h.clients, conn)
			}
		}
		h.mu.Unlock()
	}
}

func (h *Hub) add(conn *websocket.Conn) {
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) remove(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()
	conn.Close()
}

func newIPLimiter() *ipLimiter {
	ipl := &ipLimiter{
		limiters: make(map[string]*visitorLimiter),
	}
	go ipl.cleanup()
	return ipl
}

func (ipl *ipLimiter) getLimiter(ip string) *rate.Limiter {
	ipl.mu.Lock()
	defer ipl.mu.Unlock()

	v, exists := ipl.limiters[ip]
	if !exists {
		v = &visitorLimiter{
			limiter: rate.NewLimiter(10, 15),
		}
		ipl.limiters[ip] = v
	}
	v.lastSeen = time.Now()
	return v.limiter
}

func (ipl *ipLimiter) cleanup() {
	for {
		time.Sleep(3 * time.Minute)
		ipl.mu.Lock()
		for ip, v := range ipl.limiters {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(ipl.limiters, ip)
			}
		}
		ipl.mu.Unlock()
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func remoteIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		if ip, _, err := net.SplitHostPort(forwarded); err == nil {
			return ip
		}
		return forwarded
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func StartServer(db *gorm.DB, cfg config.Config) *Hub {
	router := chi.NewRouter()

	ipl := newIPLimiter()
	rateLimitMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			limiter := ipl.getLimiter(remoteIP(r))
			if !limiter.Allow() {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Setup middleware
	router.Use(middleware.Logger)
	router.Use(corsMiddleware)
	router.Use(rateLimitMiddleware)

	// setup fs folders & ws hub
	hub := newHub()
	go hub.run()

	publicFS, err := fs.Sub(app.Dist, "public")
	if err != nil {
		log.Fatal(err)
	}
	router.Mount("/public", http.StripPrefix("/public/", http.FileServer(http.FS(publicFS))))

	distFS, err := fs.Sub(app.Dist, "dist")
	if err != nil {
		log.Fatal(err)
	}

	const pingInterval = 30 * time.Second
	const connDeadline = 60 * time.Second

	// This provides a websocket that sends each repository scanned to clients
	router.Get("/ws/live", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("couldnt upgrade websocket!")
			return
		}

		conn.SetReadDeadline(time.Now().Add(connDeadline))
		conn.SetPongHandler(func(string) error {
			conn.SetReadDeadline(time.Now().Add(connDeadline))
			return nil
		})

		hub.add(conn)
		defer hub.remove(conn)

		ticker := time.NewTicker(pingInterval)
		defer ticker.Stop()

		done := make(chan struct{})
		go func() {
			defer close(done)
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					return
				}
			}
		}()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	})

	// Handles API

	InitRepositories(router, db)
	InitFindings(router, db)
	InitLeaderboard(router, db, distFS)
	InitDocumentation(router, db, distFS)

	fileServer := http.FileServer(http.FS(distFS))

	// This handles index.html and whatnot
	router.Handle("/*", fileServer)

	// This is where we serve our api & content.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.HTTP.Port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Listens for conns & serves
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()
	return hub
}
