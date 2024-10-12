package services

import (
	"btmho/app/models"
	"btmho/app/repositories"

	"errors"
)

// AuthService define a interface para o serviço de autenticacao
type AuthService interface {
	RegisterUser(user *models.Usuario) error
	Login(credentials models.Credenciais) (string, error)
	RecoverPassword(email string) error
}

// authService é a implementação concreta de AuthService
type authService struct {
	userRepo        repositories.UserRepository
	userValidator   UserValidator
	passwordService PasswordService
	tokenService    TokenService
	emailService    EmailService
}

// NewAuthService cria uma nova instância de AuthService
func NewAuthService(userRepo repositories.UserRepository, userValidator UserValidator, passwordService PasswordService, tokenService TokenService, emailService EmailService) AuthService {
	return &authService{userRepo: userRepo, userValidator: userValidator, passwordService: passwordService, tokenService: tokenService, emailService: emailService}
}

// RegisterUser orquestra o registro de um novo usuário
func (s *authService) RegisterUser(user *models.Usuario) error {
	// Valida os dados do usuário
	if err := s.userValidator.Validate(user); err != nil {
		return err
	}

	// Verifica se o usuário já existe pelo e-mail
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Faz o hash da senha
	hashedPassword, err := s.passwordService.HashPassword(user.Senha)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Senha = hashedPassword

	// Cria o usuário no repositório
	if err := s.userRepo.CreateUser(user); err != nil {
		return errors.New("error saving user")
	}

	return nil
}

// Login validates user credentials and returns a JWT token
func (s *authService) Login(credenciais models.Credenciais) (string, error) {
	// Fetch user by email
	user, err := s.userRepo.GetUserByEmail(credenciais.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Check the password
	if !s.passwordService.CheckPasswordHash(credenciais.Senha, user.Senha) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token for the authenticated user
	token, err := s.tokenService.GenerateJWT(user.Id)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}

// RecoverPassword generates a password recovery token and sends it via email
func (s *authService) RecoverPassword(email string) error {
	// Check if the user exists
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	// Generate recovery token (this method should be implemented in your token service)
	token, err := s.tokenService.GeneratePasswordRecoveryToken(email)
	if err != nil {
		return errors.New("error generating recovery token")
	}

	// Send the recovery email
	if err := s.emailService.SendRecoveryEmail(email, token); err != nil {
		return errors.New("error sending recovery email")
	}

	return nil
}
