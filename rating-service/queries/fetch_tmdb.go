package queries

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var redisClient *redis.Client

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // set to your password
		DB:       0,                // default DB
	})

	// Test the connection
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis")
}

func setDataInRedis(c *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(c.Request.Context(), key, value, expiration).Err()
}

func FetchRating(c *gin.Context) {
	log.Println("Inside FetchRating")
	id := c.Param("id")
	cacheKey := "movie:" + id

	if redisClient == nil {
		log.Println("Redis client is nil. Initializing...")
		initRedis()
	}

	ctx := c.Request.Context()
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Data found in cache")
		var movieData map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &movieData); err != nil {
			log.Printf("Failed to parse cached data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse cached data"})
			return
		}
		c.JSON(http.StatusOK, movieData)
		return
	} else if err != redis.Nil {
		log.Printf("Redis error: %v", err)
	} else {
		log.Println("Data not found in cache")
	}

	// Data not in cache or error occurred, fetch from API
	url := "https://api.themoviedb.org/3/movie/" + id + "?language=en-US"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	apiKey := os.Getenv("TMDB_API")
	if apiKey == "" {
		log.Println("API key not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not found"})
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from API"})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	if len(body) == 0 {
		log.Println("Received empty response from API")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Received empty response from API"})
		return
	}

	var movieData map[string]interface{}
	if err := json.Unmarshal(body, &movieData); err != nil {
		log.Printf("Failed to parse API response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	// Process the data if needed
	processedData := gin.H{
		"id":           movieData["id"],
		"title":        movieData["title"],
		"rating":       movieData["vote_average"],
		"vote_count":   movieData["vote_count"],
		"overview":     movieData["overview"],
		"popularity":   movieData["popularity"],
		"runtime":      movieData["runtime"],
		"release_date": movieData["release_date"],
	}

	// Cache the processed data in Redis
	jsonData, err := json.Marshal(processedData)
	if err != nil {
		log.Printf("Failed to marshal data for caching: %v", err)
	} else {
		err = setDataInRedis(c, cacheKey, jsonData, 1*time.Hour) // Cache for 1 hour
		if err != nil {
			log.Printf("Redis caching error: %v", err)
		} else {
			log.Println("Successfully cached data in Redis")
		}
	}

	c.JSON(http.StatusOK, processedData)
}
