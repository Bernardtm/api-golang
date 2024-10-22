package users

import "btmho/app/domain/address"

type User struct {
	Id              string          `json:"id,omitempty" bson:"_id,omitempty"`
	FullName        string          `json:"full_name" validate:"required"`
	Email           string          `json:"email" validate:"required,email"`
	Password        string          `json:"password" validate:"required,min=8"`
	ConfirmPassword string          `json:"confirm_password" validate:"required,min=8,eqfield=Password" bson:"-"`
	Address         address.Address `json:"address"`
}
