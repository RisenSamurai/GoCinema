package main

import (
	"GoCinema/src/lib/server/database"
	"GoCinema/src/lib/server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var client *mongo.Client

func main() {
	r := gin.Default()

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	database.Cn() //Database connection

	r.POST("/add-actor", handlers.AddActor)

	r.Run(serverAddress)
}
