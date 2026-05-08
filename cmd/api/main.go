package main

import (
	"Users/max/Tech/go-boilderplate/internal/album"
	"Users/max/Tech/go-boilderplate/internal/logger"
	albumrouter "Users/max/Tech/go-boilderplate/internal/router"
	"Users/max/Tech/go-boilderplate/internal/storage/postgres"
)

func main() {
	log := logger.New()

	database, err := postgres.NewDb(log)

	if err != nil {
		log.Error("failed to connect database",
			"error", err)
	}

	repo := album.NewRepository(database)
	service := album.NewService(repo)
	handler := album.NewHandler(service)
	router := albumrouter.NewRouter(handler, log)

	log.Info("starting server", "port", "8080")

	// Start a new server
	if err := router.Run(":8080"); err != nil {
		log.Error("server failed", "error", err)
	}
}
