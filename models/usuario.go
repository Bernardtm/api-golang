package models

type Usuario struct {
	Id             string   `json:"id,omitempty" bson:"_id,omitempty"`
	NomeCompleto   string   `json:"nome_completo" validate:"required"`
	Email          string   `json:"email" validate:"required,email"`
	Senha          string   `json:"senha" validate:"required,min=8"`
	ConfirmarSenha string   `json:"confirmar_senha" validate:"required,min=8"`
	Endereco       Endereco `json:"endereco"`
}

type PasswordRecoveryRequest struct {
	Email string `json:"email"`
}
