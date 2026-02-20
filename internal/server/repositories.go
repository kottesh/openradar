package server

import (
	"log"
	"net/http"
	"strconv"

	"openradar/internal/api"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func InitRepositories(router chi.Router, db *gorm.DB) {
	// This returns information on a repository
	router.Get("/api/repository", func(w http.ResponseWriter, r *http.Request) {
		repoUrl := r.URL.Query().Get("repo_url")
		if repoUrl == "" {
			http.Error(w, "repo_url parameter is required", http.StatusBadRequest)
			return
		}

		repositories, err := api.GetRepositoryInfo(repoUrl, db)
		if err != nil {
			log.Printf("GET /repository error: %v", err)
			http.Error(w, "failed to fetch repository", http.StatusInternalServerError)
			return
		}

		if len(repositories) == 0 {
			http.Error(w, "Repository not found", http.StatusNotFound)
			return
		}

		writeJSON(w, http.StatusOK, repositories[0])
	})

	// This gets the amount of repositories scanned.
	router.Get("/api/repositories/count", func(w http.ResponseWriter, r *http.Request) {
		count := api.GetRepositoriesCount()

		writeJSON(w, http.StatusOK, map[string]int64{"total_count": count})
	})

	// This gets the findings for a specific repository
	router.Get("/api/repository/findings", func(w http.ResponseWriter, r *http.Request) {
		repoUrl := r.URL.Query().Get("repo_url")
		if repoUrl == "" {
			http.Error(w, "repo_url parameter is required", http.StatusBadRequest)
			return
		}

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

		paginatedFindings, err := api.GetFindingsFromRepository(repoUrl, page, pageSize, db)
		if err != nil {
			log.Printf("GET /repository/findings error: %v", err)
			http.Error(w, "failed to fetch findings", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, paginatedFindings)
	})

	// This gets scanned repositories by page (pagination)
	router.Get("/api/repositories", func(w http.ResponseWriter, r *http.Request) {
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

		paginatedRepos, err := api.GetAllRepositories(page, pageSize, db)
		if err != nil {
			log.Printf("GET /repositories error: %v", err)
			http.Error(w, "failed to fetch repositories", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, paginatedRepos)
	})
}
