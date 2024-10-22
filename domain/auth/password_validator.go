package auth

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
func ValidatePassword(password, confirmPassword string) error {
	var errors []string

	if password != confirmPassword {
		errors = append(errors, "password and confirmation password do not match.")
	}

	// Check for length
	if len(password) < 8 {
		errors = append(errors, "password must be at least 8 characters long.")
	}

	// Check for at least one uppercase letter
	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		errors = append(errors, "password must contain at least one uppercase letter.")
	}

	// Check for at least one lowercase letter
	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		errors = append(errors, "password must contain at least one lowercase letter.")
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		errors = append(errors, "password must contain at least one digit.")
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password); !matched {
		errors = append(errors, "password must contain at least one special character.")
	}

	if len(errors) > 0 {
		return &PasswordValidationError{Errors: errors}
	}

	return nil
}
