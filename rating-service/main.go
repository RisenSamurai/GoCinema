package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"rating_microservice/routes"
)

func main() {

	r := gin.Default()
	log.Println("Microservice is running")
	ratingAddress := os.Getenv("RATING_ADDRESS")

	routes.SetupRatingRoutes(r)

	err := r.Run(ratingAddress)
	if err != nil {
		println("Error starting server", err)
		return
	}

}
