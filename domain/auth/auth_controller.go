package auth

import (
	"btmho/app/domain/users"
	"encoding/json"
	"net/http"
)

// AuthController handles auth-related HTTP requests
type AuthController struct {
	authService AuthService
}

// NewAuthController creates a new AuthController
func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register handles user registration
func (uc *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate user data
	if err := uc.authService.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (uc *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credenciais Credentials
	if err := json.NewDecoder(r.Body).Decode(&credenciais); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := uc.authService.Login(credenciais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// PasswordRecovery handles password recovery requests
func (uc *AuthController) PasswordRecovery(w http.ResponseWriter, r *http.Request) {
	var request PasswordRecoveryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := uc.authService.RecoverPassword(request.Email); err != nil {
		http.Error(w, "Error generating recovery token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
