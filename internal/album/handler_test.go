package album

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUpRouter(h *Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/albums", h.GetAlbums)
	r.GET("/albums/:id", h.GetAlbumsById)
	r.POST("/albums", h.CreateAlbums)

	return r
}

func TestHandler_GetAlbums(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(m *MockRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			mockSetup: func(m *MockRepository) {
				m.On("GetAll").Return([]Album{
					{
						Title:  "Album 1",
						Artist: "Artist 1",
						Price:  1.0,
						ID:     int64(1),
					},
				}, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"title":"Album 1","artist":"Artist 1","price":1.0}]`,
		},
		{
			name: "InternalServerError",
			mockSetup: func(m *MockRepository) {
				m.On("GetAll").Return([]Album{}, errors.New("error")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to fetch albums"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tc.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			router := setUpRouter(handler)

			req, err := http.NewRequest(http.MethodGet, "/albums", nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAlbumsById(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(m *MockRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			id:   "1",
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", int64(1)).Return(&Album{
					Title:  "Album 1",
					Artist: "Artist 1",
					Price:  1.0,
					ID:     1,
				}, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"Album 1","artist":"Artist 1","price":1.0}`,
		},
		{
			name:           "Invalid ID",
			id:             "a",
			mockSetup:      func(m *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid album id"}`,
		},
		{
			name: "Not Found",
			id:   "2",
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", int64(2)).Return(nil, ErrNotFound).Once()
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"album not found"}`,
		},
		{
			name: "Internal Server Error",
			id:   "1",
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", int64(1)).Return(nil, errors.New("different error")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to fetch album"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tc.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			router := setUpRouter(handler)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/albums/%s", tc.id), nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateAlbums(t *testing.T) {
	tests := []struct {
		name           string
		body           CreateAlbumRequest
		mockSetup      func(m *MockRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Created",
			body: CreateAlbumRequest{
				Title:  "Album 1",
				Artist: "Artist 1",
				Price:  1.0,
			},
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.Anything).Return(nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":0,"title":"Album 1","artist":"Artist 1","price":1.0}`,
		},
		{
			name:           "BadRequest",
			body:           CreateAlbumRequest{Title: "Album 1", Artist: "Artist 1"},
			mockSetup:      func(m *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid request body"}`,
		},
		{
			name: "InternalServerError",
			body: CreateAlbumRequest{Title: "Album 1", Artist: "Artist 1", Price: 1.0},
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.Anything).Return(errors.New("some error")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to create album"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tc.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			router := setUpRouter(handler)

			bodyBytes, _ := json.Marshal(tc.body)

			req, err := http.NewRequest(http.MethodPost, "/albums", bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
			mockRepo.AssertExpectations(t)
		})
	}
}
