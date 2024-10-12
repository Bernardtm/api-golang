package services

import (
	"btmho/app/models"
	"btmho/app/repositories"

	"errors"
)

// UserService define a interface para o serviço de usuários
type UserService interface {
	ListUsers() ([]models.Usuario, error)
}

// userService é a implementação concreta de UserService
type userService struct {
	userRepo        repositories.UserRepository
	userValidator   UserValidator
	passwordService PasswordService
	tokenService    TokenService
	emailService    EmailService
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// ListUsers retrieves all users from the repository
func (s *userService) ListUsers() ([]models.Usuario, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, errors.New("error fetching users")
	}
	return users, nil
}
