package main

import (
	"Users/max/Tech/go-boilderplate/internal/album"
	router2 "Users/max/Tech/go-boilderplate/internal/router"
	"Users/max/Tech/go-boilderplate/internal/storage/postgres"
	"log"
)

func main() {
	database, err := postgres.NewDb()

	if err != nil {
		log.Fatal(err)
	}

	repo := album.NewRepository(database)
	handler := album.NewHandler(repo)
	router := router2.NewRouter(handler)

	// Start a new server
	router.Run("localhost:8080")
}
