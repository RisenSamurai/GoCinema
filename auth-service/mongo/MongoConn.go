package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Cn() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbAddress := os.Getenv("DB_ADDRESS")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	if err != nil {
		log.Println("Error connecting to MongoDB", err)
		return nil, err
	}

	return client, nil

}
