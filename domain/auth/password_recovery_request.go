package auth

type PasswordRecoveryRequest struct {
	Email string `json:"email"`
}