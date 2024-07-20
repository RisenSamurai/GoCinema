package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	err := r.Run("localhost:8081")
	if err != nil {
		println("Error starting server", err)
		return
	}

}
