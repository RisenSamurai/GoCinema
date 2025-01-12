package main

import (
	"auth-service/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	log.Println("Auth Service is starting...")

	authAddress := os.Getenv("AUTH_ADDRESS")

	routes.SetupAuthRoutes(r)

	r.Run(authAddress)

}
