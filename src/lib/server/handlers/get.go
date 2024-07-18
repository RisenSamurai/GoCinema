package handlers

import (
	"GoCinema/src/lib/server/database"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
