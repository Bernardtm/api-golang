// routes.go
package routes

import (
	"btmho/app/controllers"
	"btmho/app/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Apply the logger middleware to all routes
	router.Use(middlewares.Logger)

	// Rotas abertas
	router.HandleFunc("/", controllers.Status).Methods("GET")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/password-recovery", controllers.PasswordRecovery).Methods("POST")

	// Rotas protegidas
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middlewares.JWTMiddleware)
	protected.HandleFunc("/users", controllers.ListUsers).Methods("GET")

	return router
}
