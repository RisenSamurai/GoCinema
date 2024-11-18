package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func FetchPageMovie(c *gin.Context) (interface{}, error) {

	ratingAddres := os.Getenv("RATING_ADDRESS")
	url := fmt.Sprintf("http://%s/fetch/movie/%s", ratingAddres, c.Param("id"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil

}

func FetchMainPageMovies() (interface{}, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")
	url := fmt.Sprintf("http://%s/fetch/main-page-movies", ratingAddress)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var data interface{}

	if !errors.Is(err, json.NewDecoder(response.Body).Decode(&data)) {
		return "", errors.New("failed to receive tmdb data")
	}

	log.Println("data: ", data)

	return data, nil
}
