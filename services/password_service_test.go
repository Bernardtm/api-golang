package services_test

import (
	"btmho/app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := services.NewPasswordService()

	// Senha a ser testada
	password := "SecureP@ssw0rd!"

	// Faz o hash da senha
	hashedPassword, err := passwordService.HashPassword(password)

	// Verifica se não ocorreu erro
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Verifica se o hash não é igual à senha original
	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := services.NewPasswordService()

	// Senha a ser testada
	password := "SecureP@ssw0rd!"

	// Faz o hash da senha
	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	// Verifica se a senha original corresponde ao hash
	isMatch := passwordService.CheckPasswordHash(password, hashedPassword)

	// Verifica se a comparação retornou verdadeiro
	assert.True(t, isMatch)
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := services.NewPasswordService()

	// Senha correta e incorreta
	password := "SecureP@ssw0rd!"
	wrongPassword := "WrongP@ssw0rd!"

	// Faz o hash da senha correta
	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	// Verifica se a senha errada não corresponde ao hash
	isMatch := passwordService.CheckPasswordHash(wrongPassword, hashedPassword)

	// Verifica se a comparação retornou falso
	assert.False(t, isMatch)
}
