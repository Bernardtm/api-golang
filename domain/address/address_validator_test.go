package address_test

import (
	clients "btmho/app/clients/address"
	"btmho/app/domain/address"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAddressClient é o mock para o AddressClient
type MockAddressClient struct {
	mock.Mock
}

// FetchCEPData simula a chamada para buscar os dados do CEP
func (m *MockAddressClient) FetchCEPData(cep string) (*clients.AddressDTO, error) {
	args := m.Called(cep)
	if args.Get(0) != nil {
		return args.Get(0).(*clients.AddressDTO), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAddressValidator_ValidateCEP_Success(t *testing.T) {
	// Mockando o client
	mockClient := new(MockAddressClient)

	// Criando os dados mockados
	mockAddress := &clients.AddressDTO{
		Cep:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockAddress, nil)

	// Criando o validador com o mock
	validator := address.NewAddressValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userAddress := address.Address{
		CEP:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Testando o cenário de sucesso
	err := validator.ValidateCEP(userAddress)
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestAddressValidator_ValidateCEP_InvalidStreet(t *testing.T) {
	// Mockando o client
	mockClient := new(MockAddressClient)

	// Criando os dados mockados
	mockAddress := &clients.AddressDTO{
		Cep:    "20070-022",
		Street: "Avenida Paulista", // Street diferente para gerar erro
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockAddress, nil)

	// Criando o validador com o mock
	validator := address.NewAddressValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userAddress := address.Address{
		CEP:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Testando o cenário com erro na street
	err := validator.ValidateCEP(userAddress)
	assert.EqualError(t, err, "invalid Street")
	mockClient.AssertExpectations(t)
}

func TestAddressValidator_ValidateCEP_InvalidCity(t *testing.T) {
	// Mockando o client
	mockClient := new(MockAddressClient)

	// Criando os dados mockados
	mockAddress := &clients.AddressDTO{
		Cep:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "São Paulo", // City diferente para gerar erro
		State:  "RJ",
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockAddress, nil)

	// Criando o validador com o mock
	validator := address.NewAddressValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userAddress := address.Address{
		CEP:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Testando o cenário com erro na city
	err := validator.ValidateCEP(userAddress)
	assert.EqualError(t, err, "invalid City")
	mockClient.AssertExpectations(t)
}

func TestAddressValidator_ValidateCEP_InvalidState(t *testing.T) {
	// Mockando o client
	mockClient := new(MockAddressClient)

	// Criando os dados mockados
	mockAddress := &clients.AddressDTO{
		Cep:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "SP", // State diferente para gerar erro
	}

	// Configurando o mock para o client
	mockClient.On("FetchCEPData", "20070-022").Return(mockAddress, nil)

	// Criando o validador com o mock
	validator := address.NewAddressValidator(mockClient)

	// Endereço do usuário que deve ser validado
	userAddress := address.Address{
		CEP:    "20070-022",
		Street: "Street Buenos Aires",
		City:   "Rio de Janeiro",
		State:  "RJ",
	}

	// Testando o cenário com erro no estado
	err := validator.ValidateCEP(userAddress)
	assert.EqualError(t, err, "invalid State")
	mockClient.AssertExpectations(t)
}

func TestAddressValidator_ValidateCEP_FetchCEPDataError(t *testing.T) {
	// Mockando o client
	mockClient := new(MockAddressClient)

	// Configurando o mock para retornar erro
	mockClient.On("FetchCEPData", "00000-000").Return(nil, errors.New("CEP not found"))

	// Criando o validador com o mock
	validator := address.NewAddressValidator(mockClient)

	// Endereço do usuário com CEP inválido
	userAddress := address.Address{
		CEP:    "00000-000",
		Street: "Street Inexistente",
		City:   "City Inexistente",
		State:  "XX",
	}

	// Testando o cenário com erro ao buscar o CEP
	err := validator.ValidateCEP(userAddress)
	assert.EqualError(t, err, "invalid CEP: CEP not found")
	mockClient.AssertExpectations(t)
}
