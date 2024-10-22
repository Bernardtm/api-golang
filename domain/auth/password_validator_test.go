package auth_test

import (
	"btmho/app/domain/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword_Success(t *testing.T) {
	// Teste de sucesso com uma password válida
	err := auth.ValidatePassword("ValidPass1!", "ValidPass1!")
	assert.NoError(t, err)
}

func TestValidatePassword_PasswordMismatch(t *testing.T) {
	// Teste onde as passwords não coincidem
	err := auth.ValidatePassword("ValidPass1!", "DifferentPass!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password and confirmation password do not match.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_LengthError(t *testing.T) {
	// Teste onde a password é muito curta
	err := auth.ValidatePassword("Short1!", "Short1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must be at least 8 characters long.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_UppercaseError(t *testing.T) {
	// Teste onde a password não contém letras maiúsculas
	err := auth.ValidatePassword("lowercase1!", "lowercase1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one uppercase letter.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_LowercaseError(t *testing.T) {
	// Teste onde a password não contém letras minúsculas
	err := auth.ValidatePassword("UPPERCASE1!", "UPPERCASE1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one lowercase letter.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_DigitError(t *testing.T) {
	// Teste onde a password não contém dígitos
	err := auth.ValidatePassword("NoDigits!", "NoDigits!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one digit.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_SpecialCharacterError(t *testing.T) {
	// Teste onde a password não contém caracteres especiais
	err := auth.ValidatePassword("NoSpecial1", "NoSpecial1")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one special character.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_MultipleErrors(t *testing.T) {
	// Teste onde múltiplas regras são violadas
	err := auth.ValidatePassword("short", "short")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must be at least 8 characters long.")
		assert.Contains(t, validationErr.Errors, "password must contain at least one uppercase letter.")
		assert.Contains(t, validationErr.Errors, "password must contain at least one digit.")
		assert.Contains(t, validationErr.Errors, "password must contain at least one special character.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}
