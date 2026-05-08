package album

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateAlbum(t *testing.T) {
	tests := []struct {
		name          string
		input         *Album
		mockSetup     func(repository *MockRepository)
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.AnythingOfType("*album.Album")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "repo error",
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.AnythingOfType("*album.Album")).Return(errors.New("repo error"))
			},
			expectedError: errors.New("repo error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			service := NewService(mockRepo)

			tc.mockSetup(mockRepo)

			err := service.CreateAlbum(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetAlbums(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockRepository)
		expectedError  error
		expectedResult []Album
	}{
		{
			name: "success",
			mockSetup: func(m *MockRepository) {
				m.On("GetAll").Return([]Album{
					{
						ID:     1,
						Title:  "title",
						Artist: "artist",
						Price:  1.1,
					},
				}, nil)
			},
			expectedError: nil,
			expectedResult: []Album{
				{
					ID:     1,
					Title:  "title",
					Artist: "artist",
					Price:  1.1,
				},
			},
		},
		{
			name: "repo error",
			mockSetup: func(m *MockRepository) {
				m.On("GetAll").Return([]Album{}, errors.New("repo error"))
			},
			expectedError:  errors.New("repo error"),
			expectedResult: []Album{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			service := NewService(mockRepo)

			tc.mockSetup(mockRepo)
			body, err := service.GetAlbums()

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, body)
			}
		})
	}
}

func TestService_GetAlbumByID(t *testing.T) {
	tests := []struct {
		name           string
		input          int64
		mockSetup      func(*MockRepository)
		expectedError  error
		expectedResult *Album
	}{
		{
			name:  "success",
			input: int64(1),
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", int64(1)).Return(&Album{
					ID:     1,
					Title:  "title",
					Artist: "artist",
					Price:  1.1,
				}, nil)
			},
			expectedError: nil,
			expectedResult: &Album{
				ID:     1,
				Title:  "title",
				Artist: "artist",
				Price:  1.1,
			},
		},
		{
			name:  "repo error",
			input: int64(1),
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", int64(1)).Return(nil, errors.New("repo error"))
			},
			expectedError:  errors.New("repo error"),
			expectedResult: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			service := NewService(mockRepo)

			tc.mockSetup(mockRepo)

			body, err := service.GetAlbumByID(tc.input)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, body)
			}
		})
	}
}
