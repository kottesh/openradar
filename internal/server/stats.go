package server

import (
	"net/http"
	"openradar/internal/api"
	"time"

	"github.com/go-chi/chi/v5"
)

type stats struct {
	Findings     int64         `json:"findings"`
	Repositories int64         `json:"repositories"`
	Uptime       time.Duration `json:"uptime"`
}

func InitStats(router chi.Router) {
	// This returns findings count, repositories count, uptime
	router.Get("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		findings := api.GetFindingsCount()
		repositories := api.GetRepositoriesCount()
		uptime := getUptime()
		writeJSON(w, http.StatusOK, stats{
			Findings:     findings,
			Repositories: repositories,
			Uptime:       uptime,
		})
	})
}
