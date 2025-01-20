package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"oms/models"
	"time"
)

func CreateProduct(product models.Product) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err

	}
	return &product, nil

}

func GetProduct(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product models.Product

	err := ProductCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	fmt.Println(product)
	if err != nil {
		return nil, err

	}
	return &product, nil

}
func FindAllProducts() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var products []models.Product
	cursor, err := ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil

}
