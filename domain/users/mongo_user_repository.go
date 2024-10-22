// mongo_user_repository.go (Implementação específica para MongoDB)
package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserRepository(client *mongo.Client) UserRepository {
	return &MongoUserRepository{client: client}
}

func (r *MongoUserRepository) CreateUser(user *User) error {
	collection := r.client.Database("myapp").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func (r *MongoUserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	collection := r.client.Database("myapp").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Return nil and no error if no document is found
			return nil, nil
		}
		// Return the error if it's not a "no documents" error
		return nil, err
	}
	return &user, err
}

func (r *MongoUserRepository) GetAllUsers() ([]UserDTO, error) {
	var users []UserDTO
	collection := r.client.Database("myapp").Collection("users")

	// Define the projection to include only the id and email fields
	projection := bson.M{
		"_id":   1, // Include the id field
		"email": 1, // Include the email field
	}

	cur, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user UserDTO
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
