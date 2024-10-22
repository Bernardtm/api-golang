package users

import (
	"errors"
)

// UserService define a interface para o serviço de usuários
type UserService interface {
	ListUsers() ([]UserDTO, error)
}

// userService é a implementação concreta de UserService
type userService struct {
	userRepo        UserRepository
}

// NewUserService cria uma nova instância de UserService
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
