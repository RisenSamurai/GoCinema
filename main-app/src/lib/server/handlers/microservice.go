package handlers

import "net/http"

func FetchRatingApi(movieID string) (interface{}, error) {

	response, err := http.Get("http://localhost:8081/")

	return nil, nil
}
