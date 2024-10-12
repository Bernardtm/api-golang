package validators

import (
	"fmt"
	"regexp"
)

// PasswordValidationError holds the errors related to password validation
type PasswordValidationError struct {
	Errors []string
}

// Error implements the error interface for PasswordValidationError
func (p *PasswordValidationError) Error() string {
	return fmt.Sprintf("Password validation errors: %v", p.Errors)
}

// ValidatePassword checks if the password meets the required criteria
func ValidatePassword(senha, confirmarSenha string) error {
	var errors []string

	if senha != confirmarSenha {
		errors = append(errors, "password and confirmation password do not match.")
	}

	// Check for length
	if len(senha) < 8 {
		errors = append(errors, "password must be at least 8 characters long.")
	}

	// Check for at least one uppercase letter
	if matched, _ := regexp.MatchString(`[A-Z]`, senha); !matched {
		errors = append(errors, "password must contain at least one uppercase letter.")
	}

	// Check for at least one lowercase letter
	if matched, _ := regexp.MatchString(`[a-z]`, senha); !matched {
		errors = append(errors, "password must contain at least one lowercase letter.")
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString(`[0-9]`, senha); !matched {
		errors = append(errors, "password must contain at least one digit.")
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, senha); !matched {
		errors = append(errors, "password must contain at least one special character.")
	}

	if len(errors) > 0 {
		return &PasswordValidationError{Errors: errors}
	}

	return nil
}
