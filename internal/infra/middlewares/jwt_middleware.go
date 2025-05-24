package middlewares

import (
	"bernardtm/backend/internal/core/auth/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	tokenService token.TokenService
}

// NewJWTMiddleware creates a new JWTMiddleware instance
func NewJWTMiddleware(tokenService token.TokenService) *JWTMiddleware {
	return &JWTMiddleware{
		tokenService: tokenService,
	}
}

// AuthMiddleware is a middleware that protects routes by verifying JWT token
func (j *JWTMiddleware) AuthMiddleware(audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from header "Authorization"
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Expected header format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		if audience == "api" {
			// Validate JWT token
			claims, err := j.tokenService.ValidateAPIToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			if claims.Audience != audience {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid audience"})
				c.Abort()
				return
			}

			// if token is valid, stores user info in context
			c.Set("ID", claims.ID)
			c.Set("playerUUID", claims.PlayerUUID)
		} else {
			// Validate JWT token
			claims, err := j.tokenService.ValidateToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			if claims.Audience != audience {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid audience"})
				c.Abort()
				return
			}

			// if token is valid, stores user info in context
			c.Set("ID", claims.ID)
		}

		// Continue execution
		c.Next()
	}
}

// ExtractUserID is a helper to obtain the userID from context
func (j *JWTMiddleware) ExtractUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		return ""
	}

	return userID.(string)
}

// ExtractEmail is a helper to obtain the email from context
func (j *JWTMiddleware) ExtractEmail(c *gin.Context) string {
	email, exists := c.Get("email")
	if !exists {
		return ""
	}

	return email.(string)
}
