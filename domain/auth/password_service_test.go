package auth_test

import (
	"btmho/app/domain/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := auth.NewPasswordService()

	// Password a ser testada
	password := "SecureP@ssw0rd!"

	// Faz o hash da password
	hashedPassword, err := passwordService.HashPassword(password)

	// Verifica se não ocorreu erro
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Verifica se o hash não é igual à password original
	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := auth.NewPasswordService()

	// Password a ser testada
	password := "SecureP@ssw0rd!"

	// Faz o hash da password
	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	// Verifica se a password original corresponde ao hash
	isMatch := passwordService.CheckPasswordHash(password, hashedPassword)

	// Verifica se a comparação retornou verdadeiro
	assert.True(t, isMatch)
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	// Cria uma nova instância do PasswordService
	passwordService := auth.NewPasswordService()

	// Password correta e incorreta
	password := "SecureP@ssw0rd!"
	wrongPassword := "WrongP@ssw0rd!"

	// Faz o hash da password correta
	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	// Verifica se a password errada não corresponde ao hash
	isMatch := passwordService.CheckPasswordHash(wrongPassword, hashedPassword)

	// Verifica se a comparação retornou falso
	assert.False(t, isMatch)
}
