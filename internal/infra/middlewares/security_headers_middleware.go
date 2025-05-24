package middlewares

import (
	"github.com/gin-gonic/gin"
)

// TODO test more before publishing this
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content Security Policy (CSP)
		// c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")

		// X-Content-Type-Options
		// c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		// X-Frame-Options
		// c.Writer.Header().Set("X-Frame-Options", "DENY")

		// Strict-Transport-Security (HSTS)
		// Somente use HSTS se sua API suporta HTTPS
		// c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		// Continue para o próximo middleware ou handler
		c.Next()
	}
}
