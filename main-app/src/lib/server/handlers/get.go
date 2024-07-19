package handlers

import (
	"GoCinema/src/lib/server/database"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

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

func (h *Handler) GetMovie(c *gin.Context) {

	id := c.Param("id")

	movie, err := h.fetchItemFromMongo(c.Request.Context(), id)
	if err != nil {
		log.Println("Error fetching movie from Mongo: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
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
