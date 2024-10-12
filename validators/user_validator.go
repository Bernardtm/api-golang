package validators

import (
	"btmho/app/models"

	"github.com/go-playground/validator/v10"
)

// ValidateUser validates a user model
func ValidateUser(user models.Usuario) error {
	validate := validator.New()
	return validate.Struct(user)
}
