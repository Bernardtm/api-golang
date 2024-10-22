package users_test

import (
	"btmho/app/domain/address"
	"btmho/app/domain/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser_Success(t *testing.T) {
	// Teste de sucesso com um usuário válido
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		Password:        "ValidPass1!",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.NoError(t, err)
}

func TestValidateUser_MissingFullName(t *testing.T) {
	// Teste onde o campo FullName está ausente
	user := users.User{
		Email:           "bernard@example.com",
		Password:        "ValidPass1!",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidEmailFormat(t *testing.T) {
	// Teste onde o campo Email está em formato inválido
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@invalid",
		Password:        "ValidPass1!",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_MissingPassword(t *testing.T) {
	// Teste onde o campo Password está ausente
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_PasswordMismatch(t *testing.T) {
	// Teste onde Password e ConfirmPassword não são iguais
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		Password:        "ValidPass1!",
		ConfirmPassword: "InvalidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidPasswordLength(t *testing.T) {
	// Teste onde a password tem menos de 8 caracteres
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		Password:        "Short1!",
		ConfirmPassword: "Short1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_MissingAddressFields(t *testing.T) {
	// Teste onde o campo Address tem campos obrigatórios ausentes
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		Password:        "ValidPass1!",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "20070022",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidCEPLength(t *testing.T) {
	// Teste onde o CEP tem tamanho incorreto
	user := users.User{
		FullName:        "Bernard Mendes",
		Email:           "bernard@example.com",
		Password:        "ValidPass1!",
		ConfirmPassword: "ValidPass1!",
		Address: address.Address{
			Street: "Street A",
			Number: "123",
			City:   "Rio de Janeiro",
			State:  "RJ",
			CEP:    "123456",
		},
	}

	err := users.ValidateUser(user)
	assert.Error(t, err)
}
