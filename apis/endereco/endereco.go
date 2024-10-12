package endereco

import (
	"btmho/app/middlewares"
	"fmt"
	"net/http"
)

var httpGet = http.Get

func ValidateCEP(cep string) (*http.Response, error) {
	url := fmt.Sprintf(middlewares.GetDotEnvVariable("ENDERECO_API")+"/%s/json/", cep)
	return httpGet(url)
}
