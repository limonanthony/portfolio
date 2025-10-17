package main

import (
	"net/http"

	"github.com/limonanthony/portfolio/internal/env"
	"github.com/limonanthony/portfolio/internal/logger"
	"github.com/limonanthony/portfolio/internal/server"
)

func main() {
	env.Load()

	logger.Info("Starting API server...")

	serverConfig := server.NewConfig()
	srv := server.NewServer(serverConfig)

	srv.Router.Use(logger.LoggingMiddleware)

	srv.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	logger.Infof("Starting API server on port %d...", serverConfig.Port)

	if err := srv.Start(); err != nil {
		logger.Panicf("Failed to start API server: %v", err)
	}
}
