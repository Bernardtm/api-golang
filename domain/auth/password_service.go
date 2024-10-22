package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// passwordService é a implementação concreta de PasswordService
type passwordService struct{}

// NewPasswordService cria uma nova instância de PasswordService
func NewPasswordService() PasswordService {
	return &passwordService{}
}

// HashPassword faz o hash da password
func (s *passwordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash verifica se a password informada corresponde ao hash
func (s *passwordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
