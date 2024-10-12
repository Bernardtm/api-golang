package main

import (
	"context"
	"log"
	"net/http"

	clients "btmho/app/clients/endereco"
	"btmho/app/config"
	"btmho/app/db"
	"btmho/app/repositories"
	"btmho/app/routes"
	"btmho/app/services"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Connect to the database
	client := db.Connect(cfg) // Call the Connect function to establish a DB connection
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Criação de dependências
	userRepo := repositories.NewMongoUserRepository(client) // Assumindo que você tenha um client de mongo configurado
	passwordService := services.NewPasswordService()
	tokenService := services.NewTokenService(cfg)
	emailService := services.NewEmailService()
  enderecoClient := clients.NewEnderecoClient(cfg)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	router := routes.SetupRoutes(userService, authService, cfg)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
