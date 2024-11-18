package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"rating_microservice/redis_lib"
	"rating_microservice/routes"
)

func main() {

	r := gin.Default()
	log.Println("Microservice is running")
	ratingAddress := os.Getenv("RATING_ADDRESS")

	redisCall, err := redis_lib.InitRedis()
	if err != nil {
		log.Fatal(err)
	}

	defer func(redisCall *redis.Client) {
		err := redisCall.Close()
		if err != nil {

		}
	}(redisCall)

	routes.SetupRatingRoutes(r)

	err = r.Run(ratingAddress)
	if err != nil {
		println("Error starting server", err)
		return
	}

}
