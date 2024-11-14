package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func FetchTmdbExtraData(c *gin.Context, apiKey string, id string) (interface{}, error) {

	url := "https://api.themoviedb.org/3/movie/" + id + "?language=en-US"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from API"})
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return "", err
	}

	if len(body) == 0 {
		log.Println("Received empty response from API")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Received empty response from API"})
		return "", err
	}

	var movieData map[string]interface{}
	if err := json.Unmarshal(body, &movieData); err != nil {
		log.Printf("Failed to parse API response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return "", err
	}

	return movieData, nil

}
