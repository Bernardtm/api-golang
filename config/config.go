package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds the application configuration values
type AppConfig struct {
	Port          string
	MongoURI      string
	MongoMinPool  uint64
	MongoMaxPool  uint64
	JWTSecret     string
}

// LoadConfig initializes the AppConfig struct with values from environment variables
func LoadConfig() (*AppConfig, error) {
  // Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables directly.")
	}

	minPool, err := parseUint(os.Getenv("MONGO_MIN_POOL"), 1)
	if err != nil {
		return nil, err
	}
	maxPool, err := parseUint(os.Getenv("MONGO_MAX_POOL"), 10)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Port:         os.Getenv("PORT"),
		MongoURI:     os.Getenv("MONGO_URI"),
		MongoMinPool: minPool,
		MongoMaxPool: maxPool,
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}, nil
}

// parseUint is a helper function to convert a string to uint64 with a fallback value
func parseUint(value string, fallback uint64) (uint64, error) {
	if value == "" {
		return fallback, nil
	}
	return strconv.ParseUint(value, 10, 64)
}
