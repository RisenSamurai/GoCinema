package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Cn() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbAddress := os.Getenv("DB_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {

		panic(err)
	}

	return client, nil

}
