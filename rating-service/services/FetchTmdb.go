package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"rating_microservice/database"
	"rating_microservice/util"
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

func FetchRating(c *gin.Context, tmbdbID string) (interface{}, error) {

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
			return "", err
		}
		return movieData, err
	} else if !errors.Is(err, redis.Nil) {
		log.Printf("Redis error: %v", err)
	} else {
		log.Println("Data not found in cache")
	}

	apiKey := os.Getenv("TMDB_API")
	if apiKey == "" {
		log.Println("API key not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not found"})
		return "", err
	}

	// Data not in cache or error occurred, fetch from API
	movieData, err := util.FetchTmdbExtraData(c, apiKey, tmbdbID)
	if err != nil {
		log.Printf("Failed to fetch movie data: %v", err)
		return "", err
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
		"revenue":      movieData["revenue"],
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

	return processedData, nil

}

func FetchMovie(c *gin.Context) (interface{}, error) {
	movieID := c.Param("id")

	movie, err := database.FetchItemFromMongo(c, movieID)
	if err != nil {
		return "Failed to fetch item from Mongo", err
	}

	ratings, err := FetchRating(c, movie.TmdbId)
	if err != nil {
		return "Failed to fetch rating", err
	}

	items := gin.H{
		"movie":   movie,
		"ratings": ratings,
	}

	log.Println("Items to Gateway: ", items)

	return items, nil

}

func FetchItems(c *gin.Context) (interface{}, error) {
	var movies []database.Movie
	var articles []database.Article

	// Fetch movies from MongoDB
	movies, err := database.FetchAnyFromMongo[database.Movie](c.Request.Context(), "Movies")
	if err != nil {
		log.Println("Error fetching movies from Mongo:", err)
		return nil, fmt.Errorf("error fetching movies from Mongo")
	}

	// Fetch articles from MongoDB
	articles, err = database.FetchAnyFromMongo[database.Article](c.Request.Context(), "Articles")
	if err != nil {
		log.Println("Error fetching articles from Mongo:", err)
		return nil, fmt.Errorf("error fetching articles from Mongo")
	}

	// Log successful fetch
	log.Println("Successfully fetched movies and articles from Mongo")

	items := gin.H{
		"movies":   movies,
		"articles": articles,
	}

	return items, nil
}
