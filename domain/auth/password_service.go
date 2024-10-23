package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// passwordService is the concrete implementation of PasswordService
type passwordService struct{}

// NewPasswordService creates a new instance of PasswordService
func NewPasswordService() PasswordService {
	return &passwordService{}
}

// HashPassword hashes the password
func (s *passwordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash verifies if the provided password matches the hash
func (s *passwordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
