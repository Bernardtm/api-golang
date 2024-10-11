package services

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("test@example.com")
	if err != nil {
		t.Fatalf("Error generating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("Expected a JWT token, got empty string")
	}
}

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	password := "password123"

	// Testa a geração do hash
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("Expected a hashed password, got empty string")
	}

	// Testa a verificação do hash
	match := CheckPasswordHash(password, hash)
	if !match {
		t.Errorf("Expected password match, got no match")
	}
}
