package services_test

import (
	"btmho/app/config"
	"btmho/app/services"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT_Success(t *testing.T) {
	// Simula o AppConfig com um JWTSecret
	appConfig := &config.AppConfig{
		JWTSecret: "mysecretkey",
	}

	// Cria o TokenService com o segredo JWT simulado
	tokenService := services.NewTokenService(appConfig)

	// Gera um token JWT
	userID := "12345"
	token, err := tokenService.GenerateJWT(userID)

	// Verifica se não ocorreu erro
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Faz o parse do token gerado para validar seu conteúdo
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appConfig.JWTSecret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Extrai as claims do token
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, userID, claims["id"])
		assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	} else {
		t.Errorf("failed to parse claims")
	}
}

func TestGeneratePasswordRecoveryToken_Success(t *testing.T) {
	// Simula o AppConfig com um JWTSecret
	appConfig := &config.AppConfig{
		JWTSecret: "mysecretkey",
	}

	// Cria o TokenService com o segredo JWT simulado
	tokenService := services.NewTokenService(appConfig)

	// Gera um token de recuperação de senha
	email := "user@example.com"
	token, err := tokenService.GeneratePasswordRecoveryToken(email)

	// Verifica se não ocorreu erro
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Faz o parse do token gerado para validar seu conteúdo
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

func TestGenerateJWT_Error_InvalidSecret(t *testing.T) {
	// Simula um AppConfig com uma chave secreta inválida (vazia)
	appConfig := &config.AppConfig{
		JWTSecret: "",
	}

	// Cria o TokenService com o segredo JWT inválido
	tokenService := services.NewTokenService(appConfig)

	// Gera um token JWT com a chave secreta inválida
	userID := "12345"
	token, err := tokenService.GenerateJWT(userID)

	// Verifica se ocorreu erro, pois a chave secreta é inválida
	assert.Error(t, err)
	assert.Empty(t, token)
}
