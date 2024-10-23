package auth_test

import (
	"btmho/app/domain/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	passwordService := auth.NewPasswordService()

	password := "SecureP@ssw0rd!"

	hashedPassword, err := passwordService.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	passwordService := auth.NewPasswordService()

	password := "SecureP@ssw0rd!"

	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	isMatch := passwordService.CheckPasswordHash(password, hashedPassword)

	assert.True(t, isMatch)
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	passwordService := auth.NewPasswordService()

	password := "SecureP@ssw0rd!"
	wrongPassword := "WrongP@ssw0rd!"

	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)

	isMatch := passwordService.CheckPasswordHash(wrongPassword, hashedPassword)

	assert.False(t, isMatch)
}
