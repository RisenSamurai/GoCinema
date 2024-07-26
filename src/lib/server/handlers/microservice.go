package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchRatingApi(movieID string) (string, error) {
	log.Println("Inside FetchRatingApi")
	response, err := http.Get("http://ratings-service:8081/get-movie-rating/" + movieID)
	if err != nil {
		log.Println("Error fetching ratings from microservice", err)
		return "", err
	}
	defer response.Body.Close()

	var result struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Println("Error decoding response", err)
		return "", err
	}

	return result.Body, nil
}
