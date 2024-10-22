package clients

// UserDTO is a Data Transfer Object for exposing user data without sensitive information
type AddressDTO struct {
	Cep    string `json:"cep"`
	Street string `json:"logradouro"`
	City   string `json:"localidade"`
	State  string `json:"estado"`
}
