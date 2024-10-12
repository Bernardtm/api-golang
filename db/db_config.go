package db

import (
	"btmho/app/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoOptions retrieves MongoDB connection options using the provided AppConfig
func GetMongoOptions(cfg *config.AppConfig) (*options.ClientOptions, error) {
	clientOptions := options.Client().
		ApplyURI(cfg.MongoURI).
		SetMinPoolSize(cfg.MongoMinPool).
		SetMaxPoolSize(cfg.MongoMaxPool).
		SetConnectTimeout(10 * time.Second)
	return clientOptions, nil
}
