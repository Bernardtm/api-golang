// app/services/user_validator.go
package services

import (
	"btmho/app/models"
	"btmho/app/validators"
)

type UserValidator interface {
	Validate(user *models.Usuario) error
}

type userValidator struct{}

// NewUserValidator cria uma nova instância de UserValidator
func NewUserValidator() UserValidator {
	return &userValidator{}
}

// Validate valida os dados do usuário
func (v *userValidator) Validate(user *models.Usuario) error {
	if err := validators.ValidateCEP(user.Endereco.CEP); err != nil {
		return err
	}
	return validators.ValidateUser(*user)
}
