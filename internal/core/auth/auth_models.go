package auth

import (
	"bernardtm/backend/internal/core/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginRequest representa a estrutura para a requisição de login
type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"Password123#"`
}

// Login2StepRequest representa a estrutura para a requisição de 2-step login
type Login2StepRequest struct {
	OTP string `json:"otp" binding:"required" example:"123456"`
}

// RegisterRequest representa os dados de requisição de registro de usuário
type RegisterRequest struct {
	FullName        string `json:"full_name" binding:"required" example:"John Doe"`
	Email           string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Password        string `json:"password" binding:"required,min=8,max=40" example:"Password123#"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,eqfield=Password" example:"Password123#"`
}

// PasswordResetRequest represents the request data to reset the password
type PasswordResetRequest struct {
	Password        string `json:"password" binding:"required,min=8,max=40" example:"Password123#"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,eqfield=Password" example:"Password123#"`
}

// RecoverPasswordRequest representa a estrutura para a recuperação de senha
type RecoverPasswordRequest struct {
	Email string `json:"email" binding:"required,email" example:"john.doe@example.com"`
}

// TokenResponse representa a resposta após um login bem-sucedido
type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token      string  `json:"token"`
	Name       string  `json:"name"`
	PlayerUUID string  `json:"player_uuid"`
	Email      string  `json:"email"`
	Avatar     *string `json:"avatar"`
}

// UserResponse representa a resposta após o registro bem-sucedido do usuário
type UserResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

// ToUser converte RegisterRequest para User
func (r *RegisterRequest) ToUser() users.Users {
	return users.Users{
		ID:              primitive.NewObjectID(), // Gera um novo ObjectID
		FullName:        r.FullName,
		Email:           r.Email,
		Password:        r.Password,
		ConfirmPassword: r.ConfirmPassword,
		CreatedAt:       0, // Data de criação será preenchida no momento do registro
		UpdatedAt:       0, // Data de atualização será preenchida no momento do registro
	}
}
