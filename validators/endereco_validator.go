package validators

import (
	"btmho/app/models"
	clients "btmho/app/clients/endereco"
	"fmt"
  "strings"
)

// EnderecoClient defines methods for interacting with the address API
type EnderecoClient interface {
	FetchCEPData(cep string) (*clients.EnderecoDTO, error)
}

// EnderecoValidator provides validation for addresses
type EnderecoValidator struct {
	client EnderecoClient
}

// NewEnderecoValidator creates a new EnderecoValidator
func NewEnderecoValidator(client EnderecoClient) *EnderecoValidator {
	return &EnderecoValidator{client: client}
}

// ValidateCEP fetches the CEP data and validates the given address
func (v *EnderecoValidator) ValidateCEP(enderecoUsuario models.Endereco) error {
	enderecoRecebido, err := v.client.FetchCEPData(enderecoUsuario.CEP)
	if err != nil {
		return fmt.Errorf("invalid CEP: %w", err)
	}

	if err := v.compareAddressFields(enderecoUsuario, *enderecoRecebido); err != nil {
		return err
	}

	return nil
}

// compareAddressFields compares the user's address with the fetched address
func (v *EnderecoValidator) compareAddressFields(enderecoUsuario models.Endereco, enderecoRecebido clients.EnderecoDTO) error {
  if !strings.EqualFold(enderecoUsuario.Rua, enderecoRecebido.Rua) {
		return fmt.Errorf("invalid Rua")
	}
	if !strings.EqualFold(enderecoUsuario.Cidade, enderecoRecebido.Cidade) {
		return fmt.Errorf("invalid Cidade")
	}
	if !strings.EqualFold(enderecoUsuario.Estado, enderecoRecebido.Estado) {
		return fmt.Errorf("invalid Estado")
	}
	return nil
}
