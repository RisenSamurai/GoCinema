package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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

func FetchTmdbPageItem(c *gin.Context) (string, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")

	url := fmt.Sprintf("http://%s/fetch/movie/%v", ratingAddress, c.Param("id"))
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("failed to receive tmdb data")
	}

	body, err := io.ReadAll(response.Body)

	return string(body), nil
}
