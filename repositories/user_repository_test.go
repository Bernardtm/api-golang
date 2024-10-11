package repositories

import (
	"btmho/app/models"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateUser(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)
	defer client.Disconnect(context.TODO())

	user := models.User{
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Address: models.Address{
			Street: "Rua Exemplo",
			Number: "123",
			City:   "Cidade",
			State:  "Estado",
			CEP:    "12345678",
		},
	}

	if err := CreateUser(&user); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)
	defer client.Disconnect(context.TODO())

	email := "test@example.com"
	user, err := GetUserByEmail(email)
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	if user.Email != email {
		t.Errorf("Expected email %v, got %v", email, user.Email)
	}
}
