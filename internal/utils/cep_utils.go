package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Address struct {
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	ZipCode      string `json:"cep"`
}

// CepUtils contém métodos para trabalhar com CEPs
type CepUtils struct{}

// NewCepUtils cria uma nova instância de CepUtils
func NewCepUtils() *CepUtils {
	return &CepUtils{}
}

// ValidateCep valida se o CEP tem o formato correto
func (cu *CepUtils) ValidateCep(cep string) (bool, error) {
	cep = strings.TrimSpace(cep)
	cep = strings.ReplaceAll(cep, "-", "")

	if len(cep) != 8 {
		return false, errors.New("CEP inválido. Deve conter 8 dígitos")
	}

	return true, nil
}

// GetAddressByCep busca o endereço através da API ViaCEP usando o CEP fornecido
func (cu *CepUtils) GetAddressByCep(cep string) (*Address, error) {
	cep = strings.TrimSpace(cep)
	cep = strings.ReplaceAll(cep, "-", "")

	// Validar o CEP antes de fazer a requisição
	isValid, err := cu.ValidateCep(cep)
	if !isValid {
		return nil, err
	}

	// Fazer a requisição à API ViaCEP
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("erro ao consultar a API de CEP")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("falha ao buscar o CEP")
	}

	var address Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, errors.New("falha ao decodificar a resposta da API de CEP")
	}

	// Verificar se o CEP é válido (ViaCEP returns erro quando o CEP não existe)
	if address.ZipCode == "" {
		return nil, errors.New("CEP não encontrado")
	}

	return &address, nil
}
