package router

import (
	"Users/max/Tech/go-boilderplate/internal/album"
	"Users/max/Tech/go-boilderplate/internal/middlewear"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *album.Handler,
	log *slog.Logger) *gin.Engine {
	router := gin.New() // Default -> default logger, New -> custom

	router.Use(gin.Recovery())

	router.Use(middlewear.RequestIDMiddleware())
	router.Use(middlewear.Logger(log))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumsById)
	router.POST("/albums", handler.CreateAlbums)

	return router
}
