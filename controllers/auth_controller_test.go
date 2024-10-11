package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/models"
)

func TestRegister(t *testing.T) {
	// Simula um novo usuário
	newUser := models.User{
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Address: models.Address{
			Street: "Rua Exemplo",
			Number: "123",
			City:   "Cidade",
			State:  "Estado",
			CEP:    "12345678",
		},
	}

	payload, _ := json.Marshal(newUser)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	// Executa a requisição
	handler.ServeHTTP(rr, req)

	// Verifica se o status é 201 (Created)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status code %v, got %v", http.StatusCreated, status)
	}

	// Verifica a resposta
	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response["message"] != "User created successfully" {
		t.Errorf("expected message 'User created successfully', got %v", response["message"])
	}
}

func TestLogin(t *testing.T) {
	// Simula credenciais de login
	credentials := models.Credentials{
		Email:    "test@example.com",
		Password: "password123",
	}

	payload, _ := json.Marshal(credentials)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	// Executa a requisição
	handler.ServeHTTP(rr, req)

	// Verifica se o status é 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	// Verifica se o token JWT foi gerado
	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response["token"] == "" {
		t.Errorf("expected JWT token, got empty string")
	}
}

func TestPasswordRecovery(t *testing.T) {
	// Simula requisição de recuperação de senha
	request := models.PasswordRecoveryRequest{
		Email: "test@example.com",
	}

	payload, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", "/password-recovery", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PasswordRecovery)

	// Executa a requisição
	handler.ServeHTTP(rr, req)

	// Verifica se o status é 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	// Como o envio de e-mail é mockado, não há verificação de resposta
}
