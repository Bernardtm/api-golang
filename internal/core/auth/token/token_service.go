package token

import (
	"bernardtm/backend/configs"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	GenerateToken(claims jwt.Claims, audience string, timeToExpire time.Duration) (string, error)
	ValidateToken(tokenStr string) (*Claims, error)
	ValidateAPIToken(tokenStr string) (*APIClaims, error)
}

type tokenService struct {
	jwtKey []byte
}

func NewTokenService(config *configs.AppConfig) TokenService {
	return &tokenService{
		jwtKey: []byte(config.JWTSecret),
	}
}

// Generate API JWT token with id, email, name, role and permissions as claims
func (s *tokenService) GenerateToken(claims jwt.Claims, audience string, timeToExpire time.Duration) (string, error) {
	expirationTime := time.Now().Add(timeToExpire) // 24 hours

	// Set audience and expiration if the claims type includes StandardClaims
	switch c := claims.(type) {
	case *jwt.StandardClaims:
		c.ExpiresAt = expirationTime.Unix()
		c.Audience = audience
	case *Claims: // Custom claim type
		c.StandardClaims.ExpiresAt = expirationTime.Unix()
		c.StandardClaims.Audience = audience
	case *APIClaims: // Another custom claim type
		c.StandardClaims.ExpiresAt = expirationTime.Unix()
		c.StandardClaims.Audience = audience
	default:
		// If claims is of an unknown type, audience and expiration are not set
		return "", fmt.Errorf("unsupported claims type")
	}

	// Creates token using HMAC signing and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtKey)
}

// Validate JWT token and return the claims if valid
func (s *tokenService) ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Validate JWT API token and return the claims if valid
func (s *tokenService) ValidateAPIToken(tokenStr string) (*APIClaims, error) {
	claims := &APIClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
