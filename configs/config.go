package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds the application configuration values
type AppConfig struct {
	AppPort              string
	CorsOrigin           string
	GinMode              string
	SwaggerHostConfig    string
	ManagementPort       string
	MongoURI             string
	MongoMinPool         uint64
	MongoMaxPool         uint64
	JWTSecret            string
	AddressAPI           string
	PostgresDSN          string
	PostgresPoolMax      int
	MAILPIT_HOST         string
	MAILPIT_PORT         int
	S3_BUCKET_NAME       string
	S3_REGION            string
	S3_ACCESS_KEY_ID     string
	S3_SECRET_ACCESS_KEY string
	MAILGUN_API_KEY      string
	MAILGUN_DOMAIN       string
	ENVIRONMENT          string
	FrontendURL          string
	POSTGRES_DSN_TEST    string
	WS_PORT              string
	REDIS_ADDRESS        string
	DOCUMENT_DB_DSN      string
}

// LoadConfig initializes the AppConfig struct with values from environment variables
func LoadConfig(envFilePath string) (*AppConfig, error) {
	// Load the .env file
	if err := godotenv.Load(envFilePath); err != nil {
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
	mailpitPort, err := strconv.Atoi(os.Getenv("MAILPIT_PORT"))
	if err != nil {
		return nil, err
	}
	postgresPoolMax, err := strconv.Atoi(os.Getenv("POSTGRES_POOL_MAX"))
	if err != nil {
		postgresPoolMax = 25
	}

	return &AppConfig{
		AppPort:              os.Getenv("APP_PORT"),
		CorsOrigin:           os.Getenv("CORS_ORIGIN"),
		GinMode:              os.Getenv("GIN_MODE"),
		SwaggerHostConfig:    os.Getenv("SWAGGER_HOST_CFG"),
		ManagementPort:       os.Getenv("MANAGEMENT_PORT"),
		MongoURI:             os.Getenv("MONGO_URI"),
		MongoMinPool:         minPool,
		MongoMaxPool:         maxPool,
		JWTSecret:            os.Getenv("JWT_SECRET"),
		PostgresDSN:          os.Getenv("POSTGRES_DSN"),
		PostgresPoolMax:      postgresPoolMax,
		MAILPIT_HOST:         os.Getenv("MAILPIT_HOST"),
		MAILPIT_PORT:         mailpitPort,
		S3_BUCKET_NAME:       os.Getenv("S3_BUCKET_NAME"),
		S3_REGION:            os.Getenv("S3_REGION"),
		S3_ACCESS_KEY_ID:     os.Getenv("S3_ACCESS_KEY_ID"),
		S3_SECRET_ACCESS_KEY: os.Getenv("S3_SECRET_ACCESS_KEY"),
		MAILGUN_DOMAIN:       os.Getenv("MAILGUN_DOMAIN"),
		MAILGUN_API_KEY:      os.Getenv("MAILGUN_API_KEY"),
		ENVIRONMENT:          os.Getenv("ENVIRONMENT"),
		FrontendURL:          os.Getenv("FRONTEND_URL"),
		POSTGRES_DSN_TEST:    os.Getenv("POSTGRES_DSN_TEST"),
		WS_PORT:              os.Getenv("WS_PORT"),
		REDIS_ADDRESS:        os.Getenv("REDIS_ADDRESS"),
		DOCUMENT_DB_DSN:      os.Getenv("DOCUMENT_DB_DSN"),
	}, nil
}

// parseUint is a helper function to convert a string to uint64 with a fallback value
func parseUint(value string, fallback uint64) (uint64, error) {
	if value == "" {
		return fallback, nil
	}
	return strconv.ParseUint(value, 10, 64)
}
