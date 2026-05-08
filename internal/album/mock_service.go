package album

import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func (m *MockService) GetAlbums() ([]Album, error) {
	args := m.Called()
	return args.Get(0).([]Album), args.Error(1)
}

func (m *MockService) GetAlbumByID(id int64) (*Album, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Album), args.Error(1)
}

func (m *MockService) CreateAlbum(*Album) error {
	args := m.Called(m)
	return args.Error(0)
}
