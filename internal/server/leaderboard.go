package server

import (
	"openradar/internal/api"

	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func InitLeaderboard(router chi.Router, db *gorm.DB, distFS fs.FS) {
	// This returns top 3 users (top 3 meaning most leaked keys from them)
	router.Get("/api/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		findings := api.GetLeaderboardData()
		writeJSON(w, http.StatusOK, findings)
	})

	// This returns the page for the leaderboard (/* requires .html in the name, hence why we do this)
	router.Get("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		f, err := distFS.Open("leaderboard.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFileFS(w, r, distFS, "leaderboard.html")
	})

}
