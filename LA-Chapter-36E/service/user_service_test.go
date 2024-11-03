package service

import (
	"be-golang-chapter-26/LA-Chapter-36E/model"
	"be-golang-chapter-26/LA-Chapter-36E/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUser(t *testing.T) {
	t.Run("user_found", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(repository.UserRepositoryMock)
		userService := NewUserService(mockRepo)

		// Define the user to be returned by the mock
		user := &model.User{ID: 1, Name: "Alice"}

		// Setup expectations
		mockRepo.On("GetUserByID", 1).Return(user, nil)

		// Call the method
		result, err := userService.GetUser(1)

		// Assert expectations
		assert.NoError(t, err)
		assert.Equal(t, user, result)

		// Verify that the expectations were met
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(repository.UserRepositoryMock)
		userService := NewUserService(mockRepo)

		// Setup expectations
		mockRepo.On("GetUserByID", 999).Return(nil, nil)

		// Call the method
		result, err := userService.GetUser(999)

		// Assert expectations
		assert.NoError(t, err)
		assert.Nil(t, result)

		// Verify that the expectations were met
		mockRepo.AssertExpectations(t)
	})
}
