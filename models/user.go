package models

type User struct {
	Id string  `json:"id,omitempty" bson:"_id,omitempty"`
	FullName string  `json:"full_name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
	Address  Address `json:"address"`
}

type Address struct {
	Street string `json:"street" validate:"required"`
	Number string `json:"number" validate:"required"`
	City   string `json:"city" validate:"required"`
	State  string `json:"state" validate:"required"`
	CEP    string `json:"cep" validate:"required,len=8"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PasswordRecoveryRequest struct {
	Email string `json:"email"`
}
