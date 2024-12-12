package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func FetchAnyFromMongo[T any](ctx context.Context, collect string) ([]T, error) {

	client, err := Cn()
	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	collection := client.Database("GoCinema").Collection(collect)

	opts := options.Find().SetLimit(10)
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []T
	if err := cursor.All(ctx, &items); err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return nil, err
	}

	return items, nil

}

/*
func FetchItemFromMongo(ctx context.Context, id string) (Movie, error) {

	var item Movies

	client, err := Cn() // Database connection

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting id to ObjectID: ", err)
		return item, err
	}

	collection := client.Database("GoCinema").Collection("Movies")

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return item, err
	}

	return item, nil
}


*/
