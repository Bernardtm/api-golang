package validators

import (
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

func ValidateUser(user models.User) error {
	return validate.Struct(user)
}

func ValidateCEP(cep string) error {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := httpGet(url)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("Invalid CEP")
	}
	return nil
}
