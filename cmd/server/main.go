package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anuzdevops/go-distroless-image-optimization/internal/config"
	"github.com/anuzdevops/go-distroless-image-optimization/internal/handlers"
	"github.com/anuzdevops/go-distroless-image-optimization/internal/web"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	dashboardHandler, err := handlers.NewDashboardHandler()
	if err != nil {
		log.Fatalf("dashboard initialization failed: %v", err)
	}

	staticFS, err := web.Static()
	if err != nil {
		log.Fatalf("static assets initialization failed: %v", err)
	}
	metricsFile := os.Getenv("METRICS_FILE")
	if metricsFile == "" {
		metricsFile = "image-metrics.json"
	}
	mux := http.NewServeMux()
	mux.Handle("/api/metrics", handlers.NewMetricsHandler(metricsFile))
	mux.Handle("/", dashboardHandler)
	mux.Handle("/health", handlers.NewHealthHandler())
	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))),
	)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("ContainerScope listening on http://localhost:%s", cfg.Port)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}

	case <-ctx.Done():
		log.Println("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)

			if closeErr := server.Close(); closeErr != nil {
				log.Printf("server close failed: %v", closeErr)
			}
		}

		if err := <-serverErrors; !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server stopped with error: %v", err)
		}

		log.Println("ContainerScope stopped")
	}
}
