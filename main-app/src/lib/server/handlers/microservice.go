package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func FetchRatingApi(c *gin.Context, movieID string) (interface{}, error) {

	response, err := http.Get("http://localhost:8081//get-movie-rating/" + movieID)
	if err != nil {
		log.Println("Error fetching ratings from microservice", err)
		c.JSON(400, gin.H{
			"message": "Error fetching ratings from microservice",
		})
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body", err)
			c.JSON(400, gin.H{
				"message": "Error closing body",
			})
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading body", err)
		c.JSON(400, gin.H{
			"message": "Error reading body",
		})

		return nil, err
	}

	c.JSON(200, gin.H{
		"body": string(body),
	})

	return nil, nil
}
