package album

import "github.com/stretchr/testify/mock"

var _ RepositoryI = (*MockRepository)(nil)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) GetAll() ([]Album, error) {
	args := mr.Called()
	return args.Get(0).([]Album), args.Error(1)
}

func (mr *MockRepository) GetByID(id int64) (*Album, error) {
	args := mr.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Album), args.Error(1)
}

func (mr *MockRepository) Create(a *Album) error {
	args := mr.Called(a)
	return args.Error(0)
}
