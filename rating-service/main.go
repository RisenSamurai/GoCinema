package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"rating-service/queries"
)

func main() {

	r := gin.Default()
	log.Println("Microservice is running")

	r.GET("/hello", func(c *gin.Context) {
		log.Println("Inside query func")
		response, err := http.Get("http://localhost:8080/hello")
		if err != nil {
			log.Println("Error while sending request to hello service")
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error while reading response body")
			return
		}
		type Message struct {
			Text string `json:"text"`
		}
		var message Message
		err = json.Unmarshal(body, &message.Text)
		log.Println("Response body:", message.Text)
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.GET("/get-movie-rating/:id", queries.FetchRating)

	err := r.Run("localhost:8081")
	if err != nil {
		println("Error starting server", err)
		return
	}

}
