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

func FetchTmdbMainPage() (interface{}, error) {

	ratingAddress := os.Getenv("RATING_ADDRESS")

	url := fmt.Sprintf("http://%s/fetch/fetch-main-page-items", ratingAddress)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("failed to receive tmdb data")
	}

	var items interface{}

	err = json.NewDecoder(response.Body).Decode(&items)

	return items, err
}

func FetchTmdbPageItem(c *gin.Context) (interface{}, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")

	url := fmt.Sprintf("http://%s/fetch/movie/%v", ratingAddress, c.Param("id"))
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var newData interface{}
	if err != json.NewDecoder(response.Body).Decode(&newData) {
		return "", errors.New("failed to receive tmdb data")
	}

	log.Println("newData: ", newData)

	return newData, err
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

	if err != json.NewDecoder(response.Body).Decode(&data) {
		return "", errors.New("failed to receive tmdb data")
	}

	log.Println("data: ", data)

	return data, nil
}
