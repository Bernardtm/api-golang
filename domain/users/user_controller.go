package users

import (
	"encoding/json"
	"net/http"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userService UserService
}

// NewUserController creates a new UserController
func NewUserController(userService UserService) *UserController {
	return &UserController{userService: userService}
}

// ListUsers retrieves all users
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.ListUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
