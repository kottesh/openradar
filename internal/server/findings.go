package server

// This initialises the APIs referenced by the frontend
// for the findings, so we can find api keys and
// whatnot.

import (
	"log"
	"net/http"
	"strconv"

	"openradar/internal/api"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func InitFindings(router chi.Router, db *gorm.DB) {
	// This returns the findings (scraped api keys)
	router.Get("/api/findings", func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page := 1
		if pageStr != "" {
			if val, err := strconv.Atoi(pageStr); err == nil && val > 0 {
				page = val
			}
		}

		pageSizeStr := r.URL.Query().Get("page_size")
		pageSize := 25
		if pageSizeStr != "" {
			if val, err := strconv.Atoi(pageSizeStr); err == nil && val > 0 && val <= 100 {
				pageSize = val
			}
		}

		provider := r.URL.Query().Get("provider")
		if provider == "" {
			provider = "*"
		}

		minAge := r.URL.Query().Get("min_age")
		if minAge == "" {
			minAge = "24h"
		}

		paginatedFindings, err := api.GetLatestFindings(
			page,
			pageSize,
			provider,
			minAge,
			db,
		)
		if err != nil {
			log.Printf("GET /findings error: %v", err)
			http.Error(w, "failed to fetch findings", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, paginatedFindings)
	})

	// This returns the count of all of the findings (scraped keys)
	router.Get("/api/findings/count", func(w http.ResponseWriter, r *http.Request) {
		count := api.GetFindingsCount()

		writeJSON(w, http.StatusOK, map[string]int64{"total_count": count})
	})
}
