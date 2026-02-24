package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"openradar/internal/config"
	"openradar/internal/db"
	"openradar/internal/jobs"
	"openradar/internal/queue"
	"openradar/internal/server"
	"openradar/internal/worker"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	queue.NewInMemoryQueue(100)

	log.Print("Initializing DB")
	database, err := db.New(cfg.Database.URL)
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}

	log.Print("Starting Server")
	hub := server.StartServer(database, cfg) // websocket

	log.Print("Starting Workers")
	for i := 0; i < cfg.Scanner.MaxConcurrentClones; i++ {
		worker.Start(ctx, cfg, database, hub)
	}

	log.Print("Initializing Jobs")
	jobs.RunJobs(jobs.JobContext{
		DB:  database,
		Cfg: cfg,
		Ctx: ctx,
	})

	// When shutting down
	<-ctx.Done()

	log.Println("Ending Process.")
}
