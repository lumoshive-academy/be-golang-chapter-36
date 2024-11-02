package service

import (
	"be-golang-chapter-26/implementation-unit-test/model"
	"be-golang-chapter-26/implementation-unit-test/repository"
)

// UserService provides methods to interact with users
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser returns a user by ID
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}