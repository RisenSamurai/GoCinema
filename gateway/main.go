package main

import (
	"GoCinema/database"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	if err != nil {
		log.Println("Error connecting to database", err)
		log.Fatal(err)
		return
	}

	r.Use(cors.Default())

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
	r.GET("/fetch-main-page-items", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/fetch-main-page-items")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error connecting to the service!", err)
		}
		defer resp.Body.Close()

		var movies []database.Movie
		err = json.NewDecoder(resp.Body).Decode(&movies)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
		}

		c.JSON(200, gin.H{
			"data": movies,
		})
	})
	r.GET("/fetch-movie/:id", func(c *gin.Context) {
		url := fmt.Sprintf("http://localhost:8082/fetch-movie/%v", c.Param("id"))
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error connecting to the service!", err)
		}

		defer resp.Body.Close()

		var movie map[string]interface{}

		err = json.NewDecoder(resp.Body).Decode(&movie)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
		}

		log.Println("I send", movie)

		c.JSON(200, gin.H{
			"movie": movie,
		})

	})

	r.Static("images/", "./static/")
	r.Run(serverAddress)
}
