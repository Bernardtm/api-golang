package auth

import (
	clients "btmho/app/clients/address"
	"btmho/app/domain/address"
	"btmho/app/domain/email"
	"btmho/app/domain/users"

	"errors"
)

// AuthService define a interface para o serviço de autenticacao
type AuthService interface {
	RegisterUser(user *users.User) error
	Login(credentials Credentials) (string, error)
	RecoverPassword(email string) error
}

// authService é a implementação concreta de AuthService
type authService struct {
	userRepo        users.UserRepository
	passwordService PasswordService
	tokenService    TokenService
	emailService    email.EmailService
	addressClient   clients.AddressClient
}

// NewAuthService cria uma nova instância de AuthService
func NewAuthService(userRepo users.UserRepository, passwordService PasswordService, tokenService TokenService, emailService email.EmailService, addressClient clients.AddressClient) AuthService {
	return &authService{userRepo: userRepo, passwordService: passwordService, tokenService: tokenService, emailService: emailService, addressClient: addressClient}
}

// RegisterUser orquestra o registro de um novo usuário
func (s *authService) RegisterUser(usuario *users.User) error {
	// Valida os dados do usuário
	if err := users.ValidateUser(*usuario); err != nil {
		return err
	}

	// Valida os dados do endereço do usuário
	if err := address.NewAddressValidator(s.addressClient).ValidateCEP(usuario.Address); err != nil {
		return err
	}

	// Validate password
	if err := ValidatePassword(usuario.Password, usuario.ConfirmPassword); err != nil {
		return err
	}

	// Verifica se o usuário já existe pelo e-mail
	existingUser, _ := s.userRepo.GetUserByEmail(usuario.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Faz o hash da password
	hashedPassword, err := s.passwordService.HashPassword(usuario.Password)
	if err != nil {
		return errors.New("error hashing password")
	}
	usuario.Password = hashedPassword

	// Cria o usuário no repositório
	if err := s.userRepo.CreateUser(usuario); err != nil {
		return errors.New("error saving user")
	}

	return nil
}

// Login validates user credentials and returns a JWT token
func (s *authService) Login(credenciais Credentials) (string, error) {
	// Fetch user by email
	user, err := s.userRepo.GetUserByEmail(credenciais.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Check the password
	if !s.passwordService.CheckPasswordHash(credenciais.Password, user.Password) {
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
