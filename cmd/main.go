package main

import (
	"log"
	"net/http"

	"btmho/app/middlewares"
	"btmho/app/routes"
)

func main() {
	// Obtém a porta das variáveis de ambiente
	port := middlewares.GetDotEnvVariable("PORT")
	if port == "" {
		port = "8080" // Valor padrão caso não exista a variável PORT no .env
	}

	router := routes.SetupRoutes()

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
