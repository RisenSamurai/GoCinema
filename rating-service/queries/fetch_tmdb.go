package queries

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var redisClient *redis.Client

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
}

type TMDBMovie struct {
	ID     string  `json:"id"`
	Rating float32 `json:"rating"`
}

func FetchRating(c *gin.Context) {

	log.Println("Inside FetchRating")
	id := c.Param("id")

	ctx := c.Request.Context()
	cachedData, err := redisClient.Get(ctx, "movie:"+id).Result()
	if err == nil {
		c.Data(200, "application/json; charset=utf-8", []byte(cachedData))
		return
	} else if !errors.Is(err, redis.Nil) {
		log.Printf("Redis error: %v", err)
	}

	url := "https://api.themoviedb.org/3/movie/" + id + "?language=en-US"

	req, _ := http.NewRequest("GET", url, nil)

	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TMDB_API")

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	err = redisClient.Set(ctx, "movie:"+id, body, 0).Err()
	if err != nil {
		log.Printf("Redis error: %v", err)

	}

	c.JSON(200, gin.H{
		"body": body,
	})

}
