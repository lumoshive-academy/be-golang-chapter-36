package repository

import (
	"golang-unit-test/model"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock implementation of UserRepository using testify/mock
type UserRepositoryMock struct {
	mock.Mock
}

// GetUserByID is a mock implementation of UserRepository.GetUserByID
func (m *UserRepositoryMock) GetUserByID(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}
