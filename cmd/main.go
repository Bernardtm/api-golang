package main

import (
	"context"
	"log"
	"net/http"

	clients "btmho/app/clients/address"
	"btmho/app/config"
	"btmho/app/db"
	"btmho/app/domain/auth"
	"btmho/app/domain/email"
	"btmho/app/domain/users"
	"btmho/app/routes"
)

// StartServer initializes and starts the server
func StartServer() error {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Connect to the database
	client := db.Connect(cfg) // Call the Connect function to establish a DB connection
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Create dependencies
	userRepo := users.NewMongoUserRepository(client)
	passwordService := auth.NewPasswordService()
	tokenService := auth.NewTokenService(cfg)
	emailService := email.NewEmailService()
	addressClient := clients.NewAddressClient(cfg)

	userService := users.NewUserService(userRepo)
	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	router := routes.SetupRoutes(userService, authService, cfg)

	log.Println("Server running on port", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, router)
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
