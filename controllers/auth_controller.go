package controllers

import (
	"btmho/app/models"
	"btmho/app/repositories"
	"btmho/app/services"
	"btmho/app/validators"
	"encoding/json"
	"log"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Valida o CEP via API externa
	if err := validators.ValidateCEP(user.Address.CEP); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Valida o restante dos dados
	if err := validators.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := services.HashPassword(user.Password)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = passwordHash

	// Verifica se o usuário já existe
	if err := repositories.CreateUser(&user); err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := repositories.GetUserByEmail(credentials.Email)
	if err != nil || !services.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateJWT(user.Id)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func PasswordRecovery(w http.ResponseWriter, r *http.Request) {
	var request models.PasswordRecoveryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Gera um token de recuperação e "envia" por e-mail (mockado)
	token, err := services.GeneratePasswordRecoveryToken(request.Email)
	if err != nil {
		http.Error(w, "Error generating recovery token", http.StatusInternalServerError)
		return
	}

	// Mock do envio de e-mail
	log.Printf("Recovery token for %s: %s\n", request.Email, token)
	w.WriteHeader(http.StatusOK)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
