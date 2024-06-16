package main

import (
	"GoCinema/src/lib/server/database"
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

	r.Run(serverAddress)
}
