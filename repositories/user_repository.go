package repositories

import (
	"btmho/app/models"
)

type UserRepository interface {
	CreateUser(user *models.Usuario) error
	GetUserByEmail(email string) (*models.Usuario, error)
	GetAllUsers() ([]models.UserDTO, error)
}
