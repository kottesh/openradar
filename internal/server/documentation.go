package server

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitDocumentation(router chi.Router, distFS fs.FS) {
	// This returns the page for the docs
	router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		f, err := distFS.Open("docs.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFileFS(w, r, distFS, "docs.html")
	})

}
