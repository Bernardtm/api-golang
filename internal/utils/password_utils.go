package utils

import (
	"errors"
	"unicode"
)

// ValidatePassword validates the strength of a password based on several criteria
func ValidatePassword(password string) []error {
	var errorsList []error

	if len(password) < 8 {
		errorsList = append(errorsList, errors.New("password must have at least 8 characters"))
	}

	if len(password) > 40 {
		errorsList = append(errorsList, errors.New("password must have less than 40 characters"))
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		errorsList = append(errorsList, errors.New("password must include at least one uppercase letter"))
	}
	if !hasLower {
		errorsList = append(errorsList, errors.New("password must include at least one lowercase letter"))
	}
	if !hasDigit {
		errorsList = append(errorsList, errors.New("password must include at least one number"))
	}
	if !hasSpecial {
		errorsList = append(errorsList, errors.New("password must include at least one special character (e.g., @, #, $, %)"))
	}

	if len(errorsList) > 0 {
		return errorsList
	}

	return nil
}
