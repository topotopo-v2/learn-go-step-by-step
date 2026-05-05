package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Starting a server (local)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
