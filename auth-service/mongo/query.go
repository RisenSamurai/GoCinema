package mongo

import (
	"auth-service/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertItemInMongo[T any](ctx context.Context, collectionName string, item T) (string, error) {
	client, err := Cn()
	if err != nil {
		log.Println("Error", err)
	}

	collection := client.Database("GoCinema").Collection(collectionName)

	_, err = collection.InsertOne(ctx, item)
	if err != nil {
		return "", err
	}

	return "Object has been inserted successfully", nil

}

func GetUserFromMongo(ctx context.Context, collectionName string, item models.User) (models.User, error) {
	client, err := Cn()
	if err != nil {
		return item, err
	}

	var user models.User

	collection := client.Database("GoCinema").Collection(collectionName)

	err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&item)
	if err != nil {
		return item, err
	}

	return item, nil

}
