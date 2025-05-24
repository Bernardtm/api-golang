package middlewares

import (
	"bernardtm/backend/internal/core/auth/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JWTQueryMiddleware struct {
	tokenService token.TokenService
}

func NewJWTQueryMiddleware(tokenService token.TokenService) *JWTQueryMiddleware {
	return &JWTQueryMiddleware{
		tokenService: tokenService,
	}
}

func (j *JWTQueryMiddleware) AuthQueryMiddleware(audience string) gin.HandlerFunc {
	return func(c *gin.Context) {

		queryToken := c.DefaultQuery("token", "")

		if queryToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		if audience == "api" {
			claims, err := j.tokenService.ValidateAPIToken(queryToken)

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

			c.Set("ID", claims.ID)
			c.Set("playerUUID", claims.PlayerUUID)
		} else {

			claims, err := j.tokenService.ValidateToken(queryToken)

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

			c.Set("ID", claims.ID)
		}

		c.Next()
	}
}
