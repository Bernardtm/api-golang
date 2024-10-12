package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"btmho/app/config"
)

// EnderecoClient holds the configuration for the address service
type EnderecoClient interface {
	FetchCEPData(cep string) (*EnderecoDTO, error)
}

type enderecoClient struct {
	apiURL     string
	httpClient *http.Client
}

// NewEnderecoClient creates a new EnderecoClient with the provided config
func NewEnderecoClient(appConfig *config.AppConfig) EnderecoClient {
	return &enderecoClient{
		apiURL:     appConfig.EnderecoAPI, // Use the API URL from appConfig
		httpClient: &http.Client{},        // You can configure this client if needed
	}
}

// ValidateCEP validates a given CEP by making an HTTP request to the address API
func (es *enderecoClient) FetchCEPData(cep string) (*EnderecoDTO, error) {
	url := fmt.Sprintf("%s/%s/json/", es.apiURL, cep)
	resp, err := es.httpClient.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid CEP")
	}

	// Parse the response
	var enderecoResponse EnderecoDTO
	if err := json.NewDecoder(resp.Body).Decode(&enderecoResponse); err != nil {
		return nil, fmt.Errorf("error parsing address data: %v", err)
	}

	// Validate address fields
	if enderecoResponse.Rua == "" || enderecoResponse.Cep == "" || enderecoResponse.Cidade == "" || enderecoResponse.Estado == "" {
		return nil, fmt.Errorf("incomplete address data")
	}

	return &enderecoResponse, nil
}
