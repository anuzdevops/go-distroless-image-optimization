package handlers

import (
	"html/template"
	"net/http"

	"github.com/anuzdevops/go-distroless-image-optimization/internal/web"
)

func NewDashboardHandler() (http.Handler, error) {
	tmpl, err := template.ParseFS(
		web.Templates(),
		"templates/index.html",
	)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusOK)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, "Failed to render dashboard", http.StatusInternalServerError)
			return
		}
	}), nil
}
