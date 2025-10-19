package main

import (
	"github.com/limonanthony/portfolio/internal/database"
	"github.com/limonanthony/portfolio/internal/database/migrations"
	"github.com/limonanthony/portfolio/internal/env"
	"github.com/limonanthony/portfolio/internal/infos"
	"github.com/limonanthony/portfolio/internal/logger"
	"github.com/limonanthony/portfolio/internal/reviews"
	"github.com/limonanthony/portfolio/internal/server"
	"gorm.io/driver/sqlite"
)

func main() {
	env.Load()

	logger.Info("Connecting to database...")
	db, err := database.NewDatabase(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		logger.Panicf("Could not connect to database: %v", err)
	}
	logger.Success("Connected to database")

	logger.Info("Running migrations...")
	if err := migrations.RunMigrations(db); err != nil {
		logger.Panicf("Could not run migrations: %v", err)
	}
	logger.Success("Migrations done...")

	logger.Info("Starting API server...")
	serverConfig := server.NewConfig()
	newServer := server.NewServer(serverConfig)
	mainRouter := newServer.Router()

	mainRouter.Use(logger.LoggingMiddleware, database.TransactionMiddleware(db))

	infos.RegisterRoutes(mainRouter)
	reviews.RegisterRoutes(mainRouter)

	logger.Infof("Starting API server on port %d...", serverConfig.Port)

	if err := newServer.Start(); err != nil {
		logger.Panicf("Failed to start API server: %v", err)
	}
}
