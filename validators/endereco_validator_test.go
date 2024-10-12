package validators_test

import (
	clients "btmho/app/clients/endereco"
	"btmho/app/models"
	"btmho/app/validators"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEnderecoClient é o mock para o EnderecoClient
type MockEnderecoClient struct {
	mock.Mock
}

// FetchCEPData simula a chamada para buscar os dados do CEP
func (m *MockEnderecoClient) FetchCEPData(cep string) (*clients.EnderecoDTO, error) {
	args := m.Called(cep)
	if args.Get(0) != nil {
		return args.Get(0).(*clients.EnderecoDTO), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestEnderecoValidator_ValidateCEP_Success(t *testing.T) {
	// Mockando o client
	mockClient := new(MockEnderecoClient)

	// Criando os dados mockados
	mockEndereco := &clients.EnderecoDTO{
		Cep:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockEndereco, nil)

	// Criando o validador com o mock
	validator := validators.NewEnderecoValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userEndereco := models.Endereco{
		CEP:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Testando o cenário de sucesso
	err := validator.ValidateCEP(userEndereco)
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestEnderecoValidator_ValidateCEP_InvalidRua(t *testing.T) {
	// Mockando o client
	mockClient := new(MockEnderecoClient)

	// Criando os dados mockados
	mockEndereco := &clients.EnderecoDTO{
		Cep:    "20070-022",
		Rua:    "Avenida Paulista", // Rua diferente para gerar erro
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockEndereco, nil)

	// Criando o validador com o mock
	validator := validators.NewEnderecoValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userEndereco := models.Endereco{
		CEP:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Testando o cenário com erro na rua
	err := validator.ValidateCEP(userEndereco)
	assert.EqualError(t, err, "invalid Rua")
	mockClient.AssertExpectations(t)
}

func TestEnderecoValidator_ValidateCEP_InvalidCidade(t *testing.T) {
	// Mockando o client
	mockClient := new(MockEnderecoClient)

	// Criando os dados mockados
	mockEndereco := &clients.EnderecoDTO{
		Cep:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "São Paulo", // Cidade diferente para gerar erro
		Estado: "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockEndereco, nil)

	// Criando o validador com o mock
	validator := validators.NewEnderecoValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userEndereco := models.Endereco{
		CEP:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Testando o cenário com erro na cidade
	err := validator.ValidateCEP(userEndereco)
	assert.EqualError(t, err, "invalid Cidade")
	mockClient.AssertExpectations(t)
}

func TestEnderecoValidator_ValidateCEP_InvalidEstado(t *testing.T) {
	// Mockando o client
	mockClient := new(MockEnderecoClient)

	// Criando os dados mockados
	mockEndereco := &clients.EnderecoDTO{
		Cep:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "SP", // Estado diferente para gerar erro
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockEndereco, nil)

	// Criando o validador com o mock
	validator := validators.NewEnderecoValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userEndereco := models.Endereco{
		CEP:    "20070-022",
		Rua:    "Rua Buenos Aires",
		Cidade: "Rio de Janeiro",
		Estado: "RJ",
	}

	// Testando o cenário com erro no estado
	err := validator.ValidateCEP(userEndereco)
	assert.EqualError(t, err, "invalid Estado")
	mockClient.AssertExpectations(t)
}

func TestEnderecoValidator_ValidateCEP_FetchCEPDataError(t *testing.T) {
	// Mockando o client
	mockClient := new(MockEnderecoClient)

	// Configurando o mock para retornar erro
	mockClient.On("FetchCEPData", "00000-000").Return(nil, errors.New("CEP not found"))

	// Criando o validador com o mock
	validator := validators.NewEnderecoValidator(mockClient)

	// Endereço do usuário com CEP inválido
	userEndereco := models.Endereco{
		CEP:    "00000-000",
		Rua:    "Rua Inexistente",
		Cidade: "Cidade Inexistente",
		Estado: "XX",
	}

	// Testando o cenário com erro ao buscar o CEP
	err := validator.ValidateCEP(userEndereco)
	assert.EqualError(t, err, "invalid CEP: CEP not found")
	mockClient.AssertExpectations(t)
}
