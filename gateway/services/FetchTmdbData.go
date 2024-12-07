package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
)

func FetchPageMovie(c *gin.Context) (interface{}, error) {

	ratingAddres := os.Getenv("RATING_ADDRESS")
	url := fmt.Sprintf("http://%s/fetch/movie/%s", ratingAddres, c.Param("id"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil

}

func FetchData(url string, resultChan chan<- interface{}, errorChan <-chan error, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		resultChan <- err
	}

	resp.Body.Close()

	var data interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		resultChan <- err
		return
	}

	resultChan <- data

}

func FetchMainPageMovies() (interface{}, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")

	moviesUrl := fmt.Sprintf("http://%s/fetch/main-page-movies", ratingAddress)
	articlesUrl := fmt.Sprintf("http://%s/fetch/articles", ratingAddress)
	seriesUrl := fmt.Sprintf("http://%s/fetch/series", ratingAddress)

	resultChan := make(chan interface{}, 3)
	errorChan := make(chan error, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	go FetchData(moviesUrl, resultChan, errorChan, &wg)
	go FetchData(articlesUrl, resultChan, errorChan, &wg)
	go FetchData(seriesUrl, resultChan, errorChan, &wg)

	wg.Wait()

	close(resultChan)
	close(errorChan)

	var results []map[string]interface{}

	for result := range resultChan {
		if resultMap, ok := result.(map[string]interface{}); ok {
			results = append(results, resultMap)
		}
	}

	if len(results) == 0 {
		return nil, errors.New("No results")
	}

	if len(results) > 1 {
		for err := range errorChan {
			log.Println("Error during data fetch", err)
		}

		return results, errors.New("Some data could not be fetched")
	}

	return results, nil
}
