package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchTmdbExtraData(apiKey, url, id string) ([]map[string]interface{}, error) {

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
	defer res.Body.Close()

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
