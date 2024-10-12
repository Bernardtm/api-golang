package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// corsMiddleware configures CORS for the application
func CORS() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"}, // Update to allow specific origins
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			ExposedHeaders:   []string{"Content-Length"},
			AllowCredentials: true,
		})

		return corsHandler.Handler(next)
	}
}
