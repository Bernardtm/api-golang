// mongo_user_repository.go (Implementação específica para MongoDB)
package repositories

import (
	"btmho/app/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserRepository(client *mongo.Client) UserRepository {
	return &MongoUserRepository{client: client}
}

func (r *MongoUserRepository) CreateUser(user *models.Usuario) error {
	collection := r.client.Database("myapp").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func (r *MongoUserRepository) GetUserByEmail(email string) (*models.Usuario, error) {
	var user models.Usuario
	collection := r.client.Database("myapp").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) GetAllUsers() ([]models.Usuario, error) {
	var users []models.Usuario
	collection := r.client.Database("myapp").Collection("users")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user models.Usuario
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
