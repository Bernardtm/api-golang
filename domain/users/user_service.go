package users

import (
	"errors"
)

// UserService defines a interface for the user service
type UserService interface {
	ListUsers() ([]UserDTO, error)
}

// userService is the UserService implementation
type userService struct {
	userRepo        UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// ListUsers retrieves all users from the repository
func (s *userService) ListUsers() ([]UserDTO, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, errors.New("error fetching users")
	}
	return users, nil
}
