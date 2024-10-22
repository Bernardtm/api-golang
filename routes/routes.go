package routes

import (
	"btmho/app/config"
	"btmho/app/domain/auth"
	"btmho/app/domain/healthcheck"
	"btmho/app/domain/users"
	"btmho/app/middlewares"
	"time"

	"github.com/gorilla/mux"
)

// RouteConfigurator is an interface to configure routes
type RouteConfigurator interface {
	Configure(router *mux.Router)
}

// SetupRoutes initializes the router and applies the middleware and routes
func SetupRoutes(userService users.UserService, authService auth.AuthService, appConfig *config.AppConfig) *mux.Router {
	router := mux.NewRouter()

	applyMiddlewares(router)
	configurePublicRoutes(router, authService)
	configureProtectedRoutes(router, userService, appConfig)

	return router
}

// applyMiddlewares applies common middlewares to the router
func applyMiddlewares(router *mux.Router) {
	router.Use(middlewares.Logger)
	router.Use(middlewares.CORS())
	router.Use(middlewares.TimeoutMiddleware(10 * time.Second))
}

// configurePublicRoutes sets up public routes
func configurePublicRoutes(router *mux.Router, authService auth.AuthService) {
	router.HandleFunc("/", healthcheck.Healthcheck).Methods("GET")
	router.HandleFunc("/auth/register", auth.NewAuthController(authService).Register).Methods("POST")
	router.HandleFunc("/auth/login", auth.NewAuthController(authService).Login).Methods("POST")
	router.HandleFunc("/auth/password-recovery", auth.NewAuthController(authService).PasswordRecovery).Methods("POST")
}

// configureProtectedRoutes sets up protected routes
func configureProtectedRoutes(router *mux.Router, userService users.UserService, appConfig *config.AppConfig) {
	protected := router.PathPrefix("/").Subrouter()

	// Initialize JWT middleware with the secret from appConfig
	jwtMiddleware := middlewares.NewJWTMiddleware(appConfig.JWTSecret)
	protected.Use(jwtMiddleware.ServeHTTP)

	protected.HandleFunc("/users", users.NewUserController(userService).ListUsers).Methods("GET")
}
