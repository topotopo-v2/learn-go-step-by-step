package router

import (
	"Users/max/Tech/go-boilderplate/internal/album"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *album.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumsById)
	router.POST("/albums", handler.CreateAlbums)

	return router
}
