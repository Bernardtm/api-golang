package configs

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectDB initializes and returns a new database connection.
func ConnectDB(config *AppConfig) (*sql.DB, error) {
	// Open the connection to PostgreSQL database
	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the PostgreSQL database: %w", err)
	}

	// Configure connection pooling
	db.SetMaxOpenConns(config.PostgresPoolMax) // Maximum open connections
	db.SetMaxIdleConns(5)                      // Minimum idle connections in the pool
	db.SetConnMaxLifetime(5 * time.Minute)     // Max time a connection can be reused

	// Use context to manage connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Ensure the database connection is alive by pinging it
	if err := db.PingContext(ctx); err != nil {
		db.Close() // Clean up resources
		return nil, fmt.Errorf("could not ping the database: %w", err)
	}

	log.Println("Connected to PostgreSQL database successfully")
	return db, nil
}
