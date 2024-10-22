package users

// UserDTO is a Data Transfer Object for exposing user data without sensitive information
type UserDTO struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Email string `json:"email"`
}
