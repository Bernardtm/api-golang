package config

import (
	"btmho/app/middlewares"
	"context"
	"log"
	"time"
	"strconv"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Connect() *mongo.Client {
	minPoolSize, err := strconv.ParseUint(middlewares.GetDotEnvVariable("MONGO_MIN_POOL"), 10, 64)
	if err != nil {
		minPoolSize = 1
	}
	maxPoolSize, err := strconv.ParseUint(middlewares.GetDotEnvVariable("MONGO_MAX_POOL"), 10, 64)
	if err != nil {
		maxPoolSize = 10
	}
	clientOptions := options.Client().
		ApplyURI(middlewares.GetDotEnvVariable("MONGO_URI")).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetConnectTimeout(10 * time.Second)
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
