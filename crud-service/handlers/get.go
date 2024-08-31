package handlers

import (
	"GoCinema/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
)

func GetItems(c *gin.Context) {

	var movies []database.Movie
	var articles []database.Article

	// Fetch movies from MongoDB
	movies, err := fetchAnyFromMongo[database.Movie](c.Request.Context(), "Movies")
	if err != nil {
		log.Println("Error fetching movies from Mongo: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching movies from Mongo"})
		return
	}

	// Fetch articles from MongoDB
	articles, err = fetchAnyFromMongo[database.Article](c.Request.Context(), "Articles")
	if err != nil {
		log.Println("Error fetching articles from Mongo: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching articles from Mongo"})
		return
	}

	// Log successful fetch
	log.Println("Successfully fetched movies and articles from Mongo")

	// Return the data to the client
	c.JSON(http.StatusOK, gin.H{
		"movies":   movies,
		"articles": articles,
	})
}

func FetchRatingApi(movieID string) (map[string]interface{}, error) {
	log.Println("Inside FetchRatingApi")
	url := fmt.Sprintf("http://localhost:8081/get-movie-rating/%s", movieID)

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching ratings from microservice: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	log.Printf("Raw response from ratings microservice: %s", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	log.Printf("Decoded response: %+v", result)

	return result, nil
}

func GetMovie(c *gin.Context) {

	log.Println("Inside GetMovie")

	id := c.Param("id")

	movie, err := fetchItemFromMongo(c.Request.Context(), id)
	ratings, err := FetchRatingApi(movie.TmdbId)
	if err != nil {
		log.Println("Error fetching movie from API: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching items from Mongo"})

	}

	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"movie":   movie,
		"ratings": ratings,
	})

}

func fetchItemFromMongo(ctx context.Context, id string) (database.Movie, error) {

	var item database.Movie

	client, err := database.Cn() // Database connection

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting id to ObjectID: ", err)
		return item, err
	}

	collection := client.Database("GoCinema").Collection("Movies")

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return item, err
	}

	return item, nil
}

func fetchAnyFromMongo[T any](ctx context.Context, collect string) ([]T, error) {

	client, err := database.Cn()
	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	collection := client.Database("GoCinema").Collection(collect)

	opts := options.Find().SetLimit(10)
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []T
	if err := cursor.All(ctx, &items); err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
		return nil, err
	}

	return items, nil

}
