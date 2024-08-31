package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func FetchTmdbMainPage() (string, error) {
	response, err := http.Get("http://localhost:8082/tmdb/fetch-main-page-items")
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("failed to receive tmdb data")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func FetchTmdbPageItem(c *gin.Context) (string, error) {
	url := fmt.Sprintf("http://localhost:8082/tmdb/%v", c.Param("id"))
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
