package users

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]UserDTO, error)
}
