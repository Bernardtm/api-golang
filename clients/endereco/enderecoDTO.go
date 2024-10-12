package clients

// UserDTO is a Data Transfer Object for exposing user data without sensitive information
type EnderecoDTO struct {
	Cep    string `json:"cep"`
	Rua string `json:"logradouro"`
  Cidade string `json:"localidade"`
  Estado string `json:"estado"`
}
