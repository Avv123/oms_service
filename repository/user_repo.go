package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"oms/config"
	"oms/models"
	"time"
)

var (
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection
)

func IntializeCollections() {
	UserCollection = config.GetCollection("users")
	ProductCollection = config.GetCollection("products")

}
func GetAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err

	}
	return users, nil

}
func CreateUser(user *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func GetByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func DeleteByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := UserCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
