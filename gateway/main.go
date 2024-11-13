package main

import (
	"GoCinema/database"
	"GoCinema/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	serverAddress := os.Getenv("GATEWAY_ADDRESS")

	if serverAddress == "" {
		log.Println("WARNING: GATEWAY_ADDRESS environment variable not set")
	}

	r.Use(cors.Default())

	routes.SetUpRoutes(r)

	r.POST("/add-actor", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/add-actor")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error connecting to the service!", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return
		}

		var message string

		if err := json.Unmarshal(body, &message); err != nil {
			log.Printf("Error decoding response: %v", err)
			return
		}

		c.JSON(200, gin.H{
			"message": message,
		})

	})
	r.POST("/add-movie", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/add-movie")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error connecting to the service!", err)
		}
		defer resp.Body.Close()

	})

	r.POST("/add-article", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/add-article")

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		if err != nil {
			log.Printf("Error creating request: %v", err)
		}

		req.Header = c.Request.Header

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error connecting to the service: %v", err)
		}

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from second API"})
			return
		}

		// Forward the response from the second API to the client
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)

	})

	r.GET("/fetch-main-page-items", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/fetch-main-page-items")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error connecting to the service!", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error connecting to the service"})
			return
		}
		defer resp.Body.Close()

		var response struct {
			Movies   []database.Movie   `json:"movies"`
			Articles []database.Article `json:"articles"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"movies":   response.Movies,
			"articles": response.Articles,
		})
	})

	r.Static("images/", "./static/")
	r.Run(serverAddress)
}
