package main

import (
	"net/http"

	"github.com/limonanthony/portfolio/internal/logger"
	"github.com/limonanthony/portfolio/internal/server"
)

func main() {
	logger.Info("Starting API server...")

	srv := server.NewServer()

	srv.Router.Use(logger.LoggingMiddleware)

	srv.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	logger.Info("Starting API server on port 8080...")

	if err := srv.Start(); err != nil {
		logger.Panicf("Failed to start API server: %v", err)
	}
}
