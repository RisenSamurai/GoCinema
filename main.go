package main

import (
	"GoCinema/src/lib/server/database"
	"GoCinema/src/lib/server/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	client, err := database.Cn() // Database connection
	if err != nil {
		log.Println("Error connecting to database", err)
		log.Fatal(err)
		return
	}

	r.Use(cors.Default())

	handler := handlers.NewHandler(client)

	r.POST("/add-actor", handler.AddActor)

	r.Static("images/", "./static/")
	r.Run(serverAddress)
}
