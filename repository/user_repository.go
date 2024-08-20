package repository

import (
	"database/sql"
	"fmt"
	"golang-unit-test/model"
)

// UserRepository defines the methods for interacting with the user repository
type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// GetUserByID is a mock implementation of UserRepository.GetUserByID
func (s *userRepository) GetUserByID(id int) (*model.User, error) {
	row := s.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
