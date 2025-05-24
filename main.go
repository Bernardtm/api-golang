// @title backend
// @description Bernardtm Backend
// @version 1.0
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide the JWT token in the format: Bearer <token>
package main

import (
	"bernardtm/backend/configs"
	docs "bernardtm/backend/docs"
	"bernardtm/backend/internal/infra/di"
	"bernardtm/backend/internal/infra/router"
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// startServer initializes and starts the server
func startServer() error {
	config, err := loadConfiguration()
	if err != nil {
		return err
	}

	db, mongoClient, err := initializeDependencies(config)
	if err != nil {
		return err
	}
	defer db.Close()

	container := di.NewContainer(db, mongoClient, config)
	mainRouter := server.SetupRouter(container, config)
	srv := createHTTPServer(config, mainRouter)
	ws := createWsServer(config, mainRouter)

	setupManagementServer(config)

	return handleGracefulShutdown(srv, ws)
}

// loadConfiguration loads the application's configuration
func loadConfiguration() (*configs.AppConfig, error) {
	cfg, err := configs.LoadConfig(".env")
	if err != nil {
		log.Printf("Failed to load configuration: %v", err)
		return nil, err
	}
	return cfg, nil
}

// initializeDependencies initializes required dependencies such as the database and mail provider
func initializeDependencies(config *configs.AppConfig) (*sql.DB, *mongo.Client, error) {
	db, err := configs.ConnectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, nil, err
	}

	mongoClient, err := configs.ConnectMongoDb(config)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, nil, err
	}

	return db, mongoClient, nil

}

// setupManagementServer sets up the management server, which serves the Swagger UI
func setupManagementServer(config *configs.AppConfig) {
	docs.SwaggerInfo.Host = config.SwaggerHostConfig

	managementRouter := gin.Default()
	managementRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go func() {
		log.Printf("Starting management server on port :%s", config.ManagementPort)
		if err := managementRouter.Run(":" + config.ManagementPort); err != nil {
			log.Fatalf("Management server failed: %v", err)
		}
	}()
}

// createHTTPServer configures and returns a new HTTP server
func createHTTPServer(config *configs.AppConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + config.AppPort,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func createWsServer(config *configs.AppConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + config.WS_PORT,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// handleGracefulShutdown listens for termination signals and shuts down the server gracefully
func handleGracefulShutdown(srv *http.Server, ws *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting ws server on port %s", ws.Addr)
		if err := ws.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
	return nil
}
