package address

type Address struct {
	Street string `json:"street" validate:"required"`
	Number string `json:"number" validate:"required"`
	City   string `json:"city" validate:"required"`
	State  string `json:"estado" validate:"required"`
	CEP    string `json:"cep" validate:"required,len=8"`
}
