package validators

import (
	"btmho/app/apis/endereco"
	"btmho/app/models"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var httpGet = http.Get

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user models.Usuario) error {
	return validate.Struct(user)
}

func ValidateCEP(cep string) error {
	resp, err := endereco.ValidateCEP(cep)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("Invalid CEP")
	}
	return nil
}
