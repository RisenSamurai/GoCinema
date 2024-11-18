package redis_lib

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var redisClient *redis.Client

func InitRedis() (*redis.Client, error) {
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

	return redisClient, nil
}

func SetDataInRedis(c *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(c.Request.Context(), key, value, expiration).Err()
}

func GetDataInRedis(c *gin.Context, key string) (string, error) {
	return redisClient.Get(c.Request.Context(), key).Result()
}
