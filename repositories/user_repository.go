package repositories

import (
	"btmho/app/config"
	"btmho/app/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {
	client = config.Connect()
}

func CreateUser(user *models.User) error {
	collection := client.Database("myapp").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	collection := client.Database("myapp").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	collection := client.Database("myapp").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user models.User
		cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}
