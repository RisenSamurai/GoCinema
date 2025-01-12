package mongo

import (
	"context"
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
