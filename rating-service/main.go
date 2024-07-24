package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rating_microservice/queries"
)

func main() {

	r := gin.Default()
	log.Println("Microservice is running")

	r.GET("/get-movie-rating/:id", queries.FetchRating)

	err := r.Run("localhost:8081")
	if err != nil {
		println("Error starting server", err)
		return
	}

}
