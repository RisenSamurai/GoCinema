package handlers

import "github.com/gin-gonic/gin"

type TMDBMovie struct {
	ID     string  `json:"id"`
	Rating float32 `json:"rating"`
}

func FetchRating(c *gin.Context) {

}
