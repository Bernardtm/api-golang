package auth

import (
	"btmho/app/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	GenerateJWT(id string) (string, error)
	GeneratePasswordRecoveryToken(email string) (string, error)
}

type tokenService struct {
	jwtKey []byte
}

func NewTokenService(appConfig *config.AppConfig) TokenService {
	return &tokenService{
		jwtKey: []byte(appConfig.JWTSecret),
	}
}

func (s *tokenService) GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(s.jwtKey)
}

func (s *tokenService) GeneratePasswordRecoveryToken(email string) (string, error) {
	return s.GenerateJWT(email) // Simula token de recuperação como JWT
}
