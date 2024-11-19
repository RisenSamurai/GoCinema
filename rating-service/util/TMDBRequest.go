package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchMovieDetails(apiKey, url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from API: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)

		return nil, err
	}

	if len(body) == 0 {
		log.Println("Received empty response from API")
		return nil, err
	}

	var decodedData map[string]interface{}

	err = json.Unmarshal(body, &decodedData)
	if err != nil {
		log.Printf("Failed to unmarshal response body: %v", err)
	}

	log.Println("Decoded data:", decodedData)
	log.Println("body:", string(body))

	return decodedData, nil

}

func FetchTmdbExtraData(apiKey, url string) ([]map[string]interface{}, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from API: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	if len(body) == 0 {
		log.Println("Received empty response from API")
		return nil, err
	}

	var response struct {
		Results []map[string]interface{} `json:"results"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Failed to parse API response: %v", err)
		return nil, err
	}

	return response.Results, nil

}
