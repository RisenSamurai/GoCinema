package handlers

import (
	"GoCinema/src/lib/server/database"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var redisClient *redis.Client

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
}

func decodeRatings(encodedRatings string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedRatings)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func (h *Handler) GetItems(c *gin.Context) {
	items, err := h.fetchItemsFromMongo(c.Request.Context())
	if err != nil {
		log.Println("Error fetching items from Mongo: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching items from Mongo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func HandleMovieRequest(c *gin.Context) {

}

func (h *Handler) GetMovie(c *gin.Context) {

	log.Println("Inside GetMovie")

	id := c.Param("id")

	ctx := c.Request.Context()
	cachedData, err := redisClient.Get(ctx, id).Bytes()
	if err == nil {
		c.Data(200, "application/json; charset=utf-8", cachedData)
	}

	movie, err := h.fetchItemFromMongo(c.Request.Context(), id)
	encodedRatings, err := FetchRatingApi(movie.TmdbId)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching items from Mongo"})

	}

	decodedRatings, err := decodeRatings(encodedRatings)
	if err != nil {
		log.Println("Error decoding movie from Mongo: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding movie from TMDB"})
	}

	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"movie":   movie,
		"ratings": decodedRatings,
	})

}

func (h *Handler) fetchItemFromMongo(ctx context.Context, id string) (database.Movie, error) {

	var item database.Movie

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting id to ObjectID: ", err)
		return item, err
	}

	collection := h.Client.Database("GoCinema").Collection("Movies")

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return item, err
	}

	return item, nil
}

func (h *Handler) fetchItemsFromMongo(ctx context.Context) ([]database.Movie, error) {
	collection := h.Client.Database("GoCinema").Collection("Movies")

	opts := options.Find().SetLimit(10)

	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("error fetching items from Mongo: %w", err)
	}
	defer cursor.Close(ctx)

	var items []database.Movie
	if err := cursor.All(ctx, &items); err != nil {
		return nil, fmt.Errorf("error decoding items from Mongo: %w", err)
	}

	return items, nil
}
