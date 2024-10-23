package users

import (
	"github.com/go-playground/validator/v10"
)

// ValidateUser validates a user model
func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}
