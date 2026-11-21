package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

func NewHealthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := HealthResponse{
			Status:    "healthy",
			Service:   "ContainerScope",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			return
		}
	})
}
