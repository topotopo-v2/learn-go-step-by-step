package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumsById)
	r.POST("/albums", postAlbums)

	return r
}

func resetAlbums() {
	albums = []album{
		{ID: "1", Title: "Test Album 1", Artist: "Artist 1", Price: 9.99},
		{ID: "2", Title: "Test Album 2", Artist: "Artist 2", Price: 19.99},
	}
}

func TestGetAlbums(t *testing.T) {
	resetAlbums()
	router := setUpRouter()

	req, _ := http.NewRequest("GET", "/albums", nil) // Why not use the error here to check?
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
}

func TestPostAlbums(t *testing.T) {
	resetAlbums()

	router := setUpRouter()

	newAlbum := album{
		ID:     "3",
		Title:  "New Album",
		Artist: "New Artist",
		Price:  29.99,
	}

	body, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newAlbum.ID, response.ID)
	assert.Equal(t, 3, len(albums))
}

func TestGetAlbumsById(t *testing.T) {
	resetAlbums()

	router := setUpRouter()

	tests := []struct {
		name   string
		id     string
		status int
	}{
		{"found", "1", http.StatusOK},
		{"not found", "10", http.StatusNotFound},
	}

	for _, tc := range tests {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%s", tc.id), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}
