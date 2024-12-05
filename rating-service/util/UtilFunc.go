package util

import (
	"fmt"
)

func FilterData[T any](data interface{}, fields []string) (map[string]interface{}, error) {
	// Ensure input is a map[string]interface{}
	rawMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input data is not a map")
	}

	// Filter the map to include only specified fields
	filteredItem := make(map[string]interface{})
	for _, field := range fields {
		if value, exists := rawMap[field]; exists {
			filteredItem[field] = value
		}
	}

	return filteredItem, nil
}

func FilterListData[T any](data interface{}, fields []string) ([]map[string]interface{}, error) {
	// Ensure input is a map[string]interface{}
	rawMap, ok := data.([]map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input data is not a map")
	}

	// Filter the map to include only specified fields
	var filteredItems []map[string]interface{}

	for _, item := range rawMap {
		filteredItem := make(map[string]interface{})

		for _, field := range fields {
			if value, exists := item[field]; exists {
				filteredItem[field] = value
			}
		}

		filteredItems = append(filteredItems, filteredItem)
	}

	return filteredItems, nil
}
