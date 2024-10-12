package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware is a middleware for validating JWT tokens
type JWTMiddleware struct {
	Secret string
}

// NewJWTMiddleware creates a new JWTMiddleware
func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{Secret: secret}
}

// ServeHTTP is the middleware function that validates JWT tokens
func (j *JWTMiddleware) ServeHTTP(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
