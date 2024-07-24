package queries

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type TMDBMovie struct {
	ID     string  `json:"id"`
	Rating float32 `json:"rating"`
}

func FetchRating(c *gin.Context) {

	log.Println("Inside FetchRating")
	id := c.Param("id")

	url := "https://api.themoviedb.org/3/movie/" + id + "?language=en-US"

	req, _ := http.NewRequest("GET", url, nil)

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TMDB_KEY")

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	c.JSON(200, gin.H{
		"body": body,
	})

}
