package models

type Endereco struct {
	Rua    string `json:"rua" validate:"required"`
	Numero string `json:"numero" validate:"required"`
	Cidade string `json:"cidade" validate:"required"`
	Estado string `json:"estado" validate:"required"`
	CEP    string `json:"cep" validate:"required,len=8"`
}
