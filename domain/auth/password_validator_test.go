package auth_test

import (
	"btmho/app/domain/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword_Success(t *testing.T) {
	err := auth.ValidatePassword("ValidPass1!", "ValidPass1!")
	assert.NoError(t, err)
}

func TestValidatePassword_PasswordMismatch(t *testing.T) {
	err := auth.ValidatePassword("ValidPass1!", "DifferentPass!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password and confirmation password do not match.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_LengthError(t *testing.T) {
	err := auth.ValidatePassword("Short1!", "Short1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must be at least 8 characters long.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_UppercaseError(t *testing.T) {
	err := auth.ValidatePassword("lowercase1!", "lowercase1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one uppercase letter.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_LowercaseError(t *testing.T) {
	err := auth.ValidatePassword("UPPERCASE1!", "UPPERCASE1!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one lowercase letter.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_DigitError(t *testing.T) {
	err := auth.ValidatePassword("NoDigits!", "NoDigits!")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one digit.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_SpecialCharacterError(t *testing.T) {
	err := auth.ValidatePassword("NoSpecial1", "NoSpecial1")
	assert.Error(t, err)

	if validationErr, ok := err.(*auth.PasswordValidationError); ok {
		assert.Contains(t, validationErr.Errors, "password must contain at least one special character.")
	} else {
		t.Errorf("expected PasswordValidationError, got %v", err)
	}
}

func TestValidatePassword_MultipleErrors(t *testing.T) {
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
