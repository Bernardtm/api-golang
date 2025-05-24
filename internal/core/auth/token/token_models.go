package token

import (
	"github.com/dgrijalva/jwt-go"
)

// Struct representing the claims of the JWT token
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type APIClaims struct {
	ID          string   `json:"id"`
	PlayerUUID  string   `json:"player_uuid"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}
