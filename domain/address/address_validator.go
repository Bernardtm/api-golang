package address

import (
	clients "btmho/app/clients/address"
	"fmt"
	"strings"
)

// AddressClient defines methods for interacting with the address API
type AddressClient interface {
	FetchCEPData(cep string) (*clients.AddressDTO, error)
}

// AddressValidator provides validation for addresses
type AddressValidator struct {
	client AddressClient
}

// NewAddressValidator creates a new AddressValidator
func NewAddressValidator(client AddressClient) *AddressValidator {
	return &AddressValidator{client: client}
}

// ValidateCEP fetches the CEP data and validates the given address
func (v *AddressValidator) ValidateCEP(addressUser Address) error {
	addressRecebido, err := v.client.FetchCEPData(addressUser.CEP)
	if err != nil {
		return fmt.Errorf("invalid CEP: %w", err)
	}

	if err := v.compareAddressFields(addressUser, *addressRecebido); err != nil {
		return err
	}

	return nil
}

// compareAddressFields compares the user's address with the fetched address
func (v *AddressValidator) compareAddressFields(addressUser Address, addressRecebido clients.AddressDTO) error {
	if !strings.EqualFold(addressUser.Street, addressRecebido.Street) {
		return fmt.Errorf("invalid Street")
	}
	if !strings.EqualFold(addressUser.City, addressRecebido.City) {
		return fmt.Errorf("invalid City")
	}
	if !strings.EqualFold(addressUser.State, addressRecebido.State) {
		return fmt.Errorf("invalid State")
	}
	return nil
}
