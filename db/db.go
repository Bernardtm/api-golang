package db

import (
	"btmho/app/config"
	"context"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Connect(cfg *config.AppConfig) *mongo.Client {
	clientOptions, err := GetMongoOptions(cfg)
	if err != nil {
		log.Fatal("⛒ GetMongoOptions Failed")
		log.Fatal(err)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
