package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`                                     // ID do MongoDB
	FullName        string             `json:"full_name" bson:"full_name" validate:"required"`                        // Nome completo (obrigatório)
	Email           string             `json:"email" bson:"email" validate:"required,email"`                          // Email (obrigatório e deve ser válido)
	Password        string             `json:"password,omitempty" bson:"password" validate:"required,min=6"`          // Senha (obrigatório, mínimo de 6 caracteres)
	ConfirmPassword string             `json:"confirm_password,omitempty" validate:"required,min=6,eqfield=Password"` // Confirmar senha (obrigatório, deve ser igual à senha)
	CreatedAt       int64              `json:"created_at" bson:"created_at"`                                          // Data de criação do usuário
	UpdatedAt       int64              `json:"updated_at" bson:"updated_at"`                                          // Data de atualização do usuário
}
