package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	log.Println("Auth Service is starting...")

	authAddres := os.Getenv("AUTH_ADDRESS")

	r.Run(":" + os.Getenv(authAddres))

}
