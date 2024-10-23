package auth

import (
	clients "btmho/app/clients/address"
	"btmho/app/domain/address"
	"btmho/app/domain/email"
	"btmho/app/domain/users"

	"errors"
)

// authService é a implementação concreta de AuthService
type AuthService interface {
	RegisterUser(user *users.User) error
	Login(credentials Credentials) (string, error)
	RecoverPassword(email string) error
}

// authService is the concrete implementation of AuthService
type authService struct {
	userRepo        users.UserRepository
	passwordService PasswordService
	tokenService    TokenService
	emailService    email.EmailService
	addressClient   clients.AddressClient
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo users.UserRepository, passwordService PasswordService, tokenService TokenService, emailService email.EmailService, addressClient clients.AddressClient) AuthService {
	return &authService{userRepo: userRepo, passwordService: passwordService, tokenService: tokenService, emailService: emailService, addressClient: addressClient}
}

// RegisterUser register a new user
func (s *authService) RegisterUser(user *users.User) error {
	if err := users.ValidateUser(user); err != nil {
		return err
	}

	if err := address.NewAddressValidator(s.addressClient).ValidateCEP(user.Address); err != nil {
		return err
	}

	if err := ValidatePassword(user.Password, user.ConfirmPassword); err != nil {
		return err
	}

	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := s.passwordService.HashPassword(user.Password)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Password = hashedPassword

	if err := s.userRepo.CreateUser(user); err != nil {
		return errors.New("error saving user")
	}

	return nil
}

// Login validates user credentials and returns a JWT token
func (s *authService) Login(credenciais Credentials) (string, error) {
	user, err := s.userRepo.GetUserByEmail(credenciais.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if !s.passwordService.CheckPasswordHash(credenciais.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := s.tokenService.GenerateJWT(user.Id)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}

// RecoverPassword generates a password recovery token and sends it via email
func (s *authService) RecoverPassword(email string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	token, err := s.tokenService.GeneratePasswordRecoveryToken(email)
	if err != nil {
		return errors.New("error generating recovery token")
	}

	if err := s.emailService.SendRecoveryEmail(email, token); err != nil {
		return errors.New("error sending recovery email")
	}

	return nil
}
