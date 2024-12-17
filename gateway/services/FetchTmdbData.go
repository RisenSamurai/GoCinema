package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"sync"
)

type Movie struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}

type MainPageData struct {
	Movies   []Movie                  `json:"movies"`
	Articles []map[string]interface{} `json:"articles"`
	Series   []map[string]interface{} `json:"series"`
}

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

func FetchData(url string, resultChan chan<- interface{}, errorChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done() // This ensures wg.Done() is called no matter what

	resp, err := http.Get(url)
	if err != nil {
		errorChan <- fmt.Errorf("error fetching URL %s: %w", url, err)
		return
	}
	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		errorChan <- fmt.Errorf("error decoding response from URL %s: %w", url, err)
		return
	}

	resultChan <- data
}

func FetchMainPageMovies() (MainPageData, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")

	moviesUrl := fmt.Sprintf("http://%s/fetch/main-page-movies", ratingAddress)
	articlesUrl := fmt.Sprintf("http://%s/fetch/articles", ratingAddress)
	seriesUrl := fmt.Sprintf("http://%s/fetch/main-page-series", ratingAddress)

	resultChan := make(chan interface{}, 3)
	errorChan := make(chan error, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	go FetchData(moviesUrl, resultChan, errorChan, &wg)
	go FetchData(articlesUrl, resultChan, errorChan, &wg)
	go FetchData(seriesUrl, resultChan, errorChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	var mainPageData MainPageData

	// Collect results and errors concurrently
	for result := range resultChan {
		if resultMap, ok := result.(map[string]interface{}); ok {
			if rawMovies, ok := resultMap["movies"]; ok {
				jsonBytes, err := json.Marshal(rawMovies)
				if err != nil {
					log.Println("Failed to marshal movies to JSON", err)
				}
				var movies []Movie
				err = json.Unmarshal(jsonBytes, &movies)
				if err != nil {
					log.Println("Failed to unmarshal movies JSON into struct", err)
				}
				mainPageData.Movies = append(mainPageData.Movies, movies...)
			}

			if rawArticles, ok := resultMap["articles"]; ok {
				if articles, ok := rawArticles.([]interface{}); ok {
					for _, article := range articles {
						if articleMap, ok := article.(map[string]interface{}); ok {
							mainPageData.Articles = append(mainPageData.Articles, articleMap)
						}
					}
				}
			}

			if rawSeries, ok := resultMap["series"]; ok {
				if series, ok := rawSeries.([]interface{}); ok {
					for _, s := range series {
						if seriesMap, ok := s.(map[string]interface{}); ok {
							mainPageData.Series = append(mainPageData.Series, seriesMap)
						}
					}
				}
			}
		}
	}

	for err := range errorChan {
		log.Println("Error during data fetch:", err)
	}

	return mainPageData, nil
}

func FetchSeriesPage(c *gin.Context) (interface{}, error) {

	ratingAddress := os.Getenv("RATING_ADDRESS")

	url := fmt.Sprintf("http://%s/fetch/series/%s", ratingAddress, c.Param("id"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func FetchArticlePage(c *gin.Context) (interface{}, error) {
	ratingAddress := os.Getenv("RATING_ADDRESS")

	resp, err := http.Get(fmt.Sprintf("http://%s/fetch/article/%s", ratingAddress, c.Param("id")))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
