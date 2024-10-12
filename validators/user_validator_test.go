package validators_test

import (
	"btmho/app/models"
	"btmho/app/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser_Success(t *testing.T) {
	// Teste de sucesso com um usuário válido
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.NoError(t, err)
}

func TestValidateUser_MissingNomeCompleto(t *testing.T) {
	// Teste onde o campo NomeCompleto está ausente
	user := models.Usuario{
		Email:          "bernard@example.com",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidEmailFormat(t *testing.T) {
	// Teste onde o campo Email está em formato inválido
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@invalid",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_MissingPassword(t *testing.T) {
	// Teste onde o campo Senha está ausente
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_PasswordMismatch(t *testing.T) {
	// Teste onde Senha e ConfirmarSenha não são iguais
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "InvalidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidPasswordLength(t *testing.T) {
	// Teste onde a senha tem menos de 8 caracteres
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		Senha:          "Short1!",
		ConfirmarSenha: "Short1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_MissingEnderecoFields(t *testing.T) {
	// Teste onde o campo Endereco tem campos obrigatórios ausentes
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "20070022",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}

func TestValidateUser_InvalidCEPLength(t *testing.T) {
	// Teste onde o CEP tem tamanho incorreto
	user := models.Usuario{
		NomeCompleto:   "Bernard Mendes",
		Email:          "bernard@example.com",
		Senha:          "ValidPass1!",
		ConfirmarSenha: "ValidPass1!",
		Endereco: models.Endereco{
			Rua:    "Rua A",
			Numero: "123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
			CEP:    "123456",
		},
	}

	err := validators.ValidateUser(user)
	assert.Error(t, err)
}
