package routes

import (
	"btmho/app/config"
	"btmho/app/controllers"
	"btmho/app/middlewares"
	"btmho/app/services"

	"github.com/gorilla/mux"
)

// RouteConfigurator is an interface to configure routes
type RouteConfigurator interface {
	Configure(router *mux.Router)
}

// SetupRoutes initializes the router and applies the middleware and routes
func SetupRoutes(userService services.UserService, authService services.AuthService, appConfig *config.AppConfig) *mux.Router {
	router := mux.NewRouter()

	// Apply common middlewares
	applyMiddlewares(router)

	// Configure routes
	configurePublicRoutes(router, authService)
	configureProtectedRoutes(router, userService, appConfig)

	return router
}

// applyMiddlewares applies common middlewares to the router
func applyMiddlewares(router *mux.Router) {
	router.Use(middlewares.Logger)
	router.Use(middlewares.TimeoutMiddleware(10))
}

// configurePublicRoutes sets up public routes
func configurePublicRoutes(router *mux.Router, authService services.AuthService) {
	router.HandleFunc("/", controllers.Healthcheck).Methods("GET")
	router.HandleFunc("/register", controllers.NewAuthController(authService).Register).Methods("POST")
	router.HandleFunc("/login", controllers.NewAuthController(authService).Login).Methods("POST")
	router.HandleFunc("/password-recovery", controllers.NewAuthController(authService).PasswordRecovery).Methods("POST")
}

// configureProtectedRoutes sets up protected routes
func configureProtectedRoutes(router *mux.Router, userService services.UserService, appConfig *config.AppConfig) {
	protected := router.PathPrefix("/").Subrouter()

	// Initialize JWT middleware with the secret from appConfig
	jwtMiddleware := middlewares.NewJWTMiddleware(appConfig.JWTSecret)
	protected.Use(jwtMiddleware.ServeHTTP)

	protected.HandleFunc("/users", controllers.NewUserController(userService).ListUsers).Methods("GET")
}
