package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchRatingApi(movieID string) (map[string]interface{}, error) {
	log.Println("Inside FetchRatingApi")
	url := fmt.Sprintf("http://localhost:8081/get-movie-rating/%s", movieID)

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching ratings from microservice: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	log.Printf("Raw response from ratings microservice: %s", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	log.Printf("Decoded response: %+v", result)

	return result, nil
}
