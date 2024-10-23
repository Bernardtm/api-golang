package auth_test

import (
	"btmho/app/config"
	"btmho/app/domain/auth"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT_Success(t *testing.T) {
	appConfig := &config.AppConfig{
		JWTSecret: "mysecretkey",
	}

	tokenService := auth.NewTokenService(appConfig)

	userID := "12345"
	token, err := tokenService.GenerateJWT(userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appConfig.JWTSecret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, userID, claims["id"])
		assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	} else {
		t.Errorf("failed to parse claims")
	}
}

func TestGeneratePasswordRecoveryToken_Success(t *testing.T) {
	appConfig := &config.AppConfig{
		JWTSecret: "mysecretkey",
	}

	tokenService := auth.NewTokenService(appConfig)

	email := "user@example.com"
	token, err := tokenService.GeneratePasswordRecoveryToken(email)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appConfig.JWTSecret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Extrai as claims do token
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, email, claims["id"])
		assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	} else {
		t.Errorf("failed to parse claims")
	}
}
