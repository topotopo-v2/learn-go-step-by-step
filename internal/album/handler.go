package album

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("Album not found")

type Handler struct {
	repo RepositoryI
}

type CreateAlbumRequest struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler(repo RepositoryI) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to fetch albums"})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *Handler) CreateAlbums(c *gin.Context) {
	var req CreateAlbumRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	album := &Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	if err := h.repo.Create(album); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, album)
}

func (h *Handler) GetAlbumsById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid album id"})
		return
	}

	album, err := h.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "album not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to fetch album"})
		return
	}

	c.JSON(http.StatusOK, album)
}
