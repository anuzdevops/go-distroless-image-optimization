package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func NewMetricsHandler(filePath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Metrics are not available. Generate the image report first.",
				http.StatusServiceUnavailable)
			return
		}

		if !json.Valid(data) {
			http.Error(w, "Invalid metrics JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	})
}
