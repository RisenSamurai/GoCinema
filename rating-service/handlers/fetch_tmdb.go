package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
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

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiMGI3N2NhNzJlYzc5MzYzMGE2MTY2NTA0ZWQ3OWQ0OSIsIm5iZiI6MTcyMTU3MDE1Mi4yMDgzMjgsInN1YiI6IjY2OWI4MWZiY2EwMWNjY2VkNjY4MjYxYyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.rMywsmtSLAmQ-q_n5Wk7vWxPK216Aozlw6P1v8565SA")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
