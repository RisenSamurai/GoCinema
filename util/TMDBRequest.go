package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func FetchTmdbExtraData(c *gin.Context, apiKey, url, id string) (interface{}, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from API: %v", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}

	if len(body) == 0 {
		log.Println("Received empty response from API")
		return "", err
	}

	var movieData map[string]interface{}
	if err := json.Unmarshal(body, &movieData); err != nil {
		log.Printf("Failed to parse API response: %v", err)
		return "", err
	}

	return movieData, nil

}
