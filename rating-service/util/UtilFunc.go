package util

import (
	"encoding/json"
	"fmt"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ParseGenres[T any](data interface{}) ([]T, error) {
	// Assert that data is a slice of interfaces
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %v", err)
	}

	// Unmarshal the JSON data into a slice of the target struct type
	var result []T
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal into target struct slice: %v", err)
	}

	return result, nil

}
