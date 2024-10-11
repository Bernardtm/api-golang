package validators

import (
	"btmho/app/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateUser(t *testing.T) {
	// Cenário de sucesso: Usuário válido
	validUser := models.User{
		FullName: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Address: models.Address{
			Street: "Rua Teste",
			Number: "123",
			City:   "Cidade Teste",
			State:  "Estado Teste",
			CEP:    "12345678",
		},
	}

	err := ValidateUser(validUser)
	if err != nil {
		t.Errorf("Expected no validation error, got %v", err)
	}

	// Cenário de falha: Usuário com email inválido
	invalidUser := models.User{
		FullName: "Jane Doe",
		Email:    "invalid-email",
		Password: "password123",
		Address: models.Address{
			Street: "Rua Teste",
			Number: "123",
			City:   "Cidade Teste",
			State:  "Estado Teste",
			CEP:    "12345678",
		},
	}

	err = ValidateUser(invalidUser)
	if err == nil {
		t.Errorf("Expected validation error for invalid email, got none")
	}
}

func TestValidateCEP(t *testing.T) {
	// Mock da função http.Get para testar a validação de CEP
	mockValidCEPResponse := `{
		"cep": "12345678",
		"logradouro": "Rua Teste",
		"complemento": "",
		"bairro": "Bairro Teste",
		"localidade": "Cidade Teste",
		"uf": "Estado Teste"
	}`

	// Mock para simular uma resposta válida da API de CEP
	httpGet = func(url string) (*http.Response, error) {
		if url == "https://viacep.com.br/ws/12345678/json/" {
			w := httptest.NewRecorder()
			w.WriteHeader(http.StatusOK)
			w.WriteString(mockValidCEPResponse)
			return w.Result(), nil
		}
		// Simula um CEP inválido
		w := httptest.NewRecorder()
		w.WriteHeader(http.StatusNotFound)
		return w.Result(), nil
	}

	// Teste com CEP válido
	err := ValidateCEP("12345678")
	if err != nil {
		t.Errorf("Expected no error for valid CEP, got %v", err)
	}

	// Teste com CEP inválido
	err = ValidateCEP("99999999")
	if err == nil {
		t.Errorf("Expected error for invalid CEP, got none")
	}
}
