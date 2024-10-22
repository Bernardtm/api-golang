package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"btmho/app/config"
)

// AddressClient holds the configuration for the address service
type AddressClient interface {
	FetchCEPData(cep string) (*AddressDTO, error)
}

type addressClient struct {
	apiURL     string
	httpClient *http.Client
}

// NewAddressClient creates a new AddressClient with the provided config
func NewAddressClient(appConfig *config.AppConfig) AddressClient {
	return &addressClient{
		apiURL:     appConfig.AddressAPI, // Use the API URL from appConfig
		httpClient: &http.Client{},       // You can configure this client if needed
	}
}

// ValidateCEP validates a given CEP by making an HTTP request to the address API
func (es *addressClient) FetchCEPData(cep string) (*AddressDTO, error) {
	url := fmt.Sprintf("%s/%s/json/", es.apiURL, cep)
	resp, err := es.httpClient.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid CEP")
	}

	// Parse the response
	var addressResponse AddressDTO
	if err := json.NewDecoder(resp.Body).Decode(&addressResponse); err != nil {
		return nil, fmt.Errorf("error parsing address data: %v", err)
	}

	// Validate address fields
	if addressResponse.Street == "" || addressResponse.Cep == "" || addressResponse.City == "" || addressResponse.State == "" {
		return nil, fmt.Errorf("incomplete address data")
	}

	return &addressResponse, nil
}
