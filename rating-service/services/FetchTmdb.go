package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"rating_microservice/database"
	"rating_microservice/redis_lib"
	"rating_microservice/util"
	"time"
)

func FetchMoviePage(c *gin.Context) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=en-US", c.Param("id"))

	cacheKey := "movie:" + c.Param("id")

	cachedData, err := redis_lib.GetDataInRedis(c, cacheKey)
	if err == redis.Nil {
		log.Println("FetchMoviePage: cache key not found. Fetching from TMDB...")

		apiKey := os.Getenv("TMDB_API")

		data, err := util.FetchMovieDetails(apiKey, url)
		if err != nil {
			log.Println("FetchMoviePage: error fetching movie details: ", err)
		}

		log.Println("FetchMoviePage: fetching movie details", data)

		fieldsToKeep := []string{"id", "title", "budget", "original_language", "genres",
			"production_companies", "origin_country", "runtime", "revenue", "overview",
			"popularity", "vote_average", "vote_count", "poster_path", "release_date",
		}

		filtered, err := util.FilterData[interface{}](data, fieldsToKeep)
		if err != nil {
			log.Println("FetchMoviePage: error filtering data: ", err)
		}

		log.Println("FetchMoviePage: filtered data: ", filtered)

		jsonData, err := json.Marshal(filtered)
		if err != nil {
			log.Println("FetchMoviePage: error marshalling data: ", err)
		}

		err = redis_lib.SetDataInRedis(c, cacheKey, jsonData, 1*time.Hour)
		if err != nil {
			log.Println("FetchMoviePage: error caching movie details: ", err)
		}

		return filtered, nil

	}

	var cachedMovie map[string]interface{}

	err = json.Unmarshal([]byte(cachedData), &cachedMovie)
	if err != nil {
		log.Println("FetchMoviePage: error unmarshalling movie details: ", err)
	}

	return cachedMovie, nil

}

func FetchMainPageMovies(c *gin.Context) ([]database.MainPageMovie, error) {

	url := "https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false" +
		"&language=en-US&page=1&sort_by=popularity.desc"

	cacheKey := "main_page_movies"
	cachedData, err := redis_lib.GetDataInRedis(c, cacheKey)
	if err == redis.Nil {
		log.Println("Cache miss for main page movies. Fetching from TMDB...")

		apiKey := os.Getenv("TMDB_API")

		data, err := util.FetchTmdbExtraData(apiKey, url)
		if err != nil {
			log.Println("Error fetching movies from TMDB:", err)
		}

		var filteredMovies []database.MainPageMovie
		for _, movie := range data {
			filteredMovies = append(filteredMovies, database.MainPageMovie{
				Id:         int(movie["id"].(float64)),
				PosterPath: movie["poster_path"].(string),
				Title:      movie["title"].(string),
			})
		}

		jsonData, err := json.Marshal(filteredMovies)
		if err != nil {
			log.Println("Error marshalling movies:", err)
			return nil, err
		}

		err = redis_lib.SetDataInRedis(c, cacheKey, jsonData, 1*time.Hour)
		if err != nil {
			log.Println("Error caching movies:", err)
			return nil, err
		}

		return filteredMovies, nil

	}

	log.Println("Cache hit for main page movies.")

	var cachedMovies []database.MainPageMovie

	err = json.Unmarshal([]byte(cachedData), &cachedMovies)
	if err != nil {
		log.Println("Error unmarshalling movies:", err)
	}

	return cachedMovies, nil

}

func FetchMainPageSeries(c *gin.Context) ([]map[string]interface{}, error) {
	url := "https://api.themoviedb.org/3/tv/latest"
	cacheKey := "main_page_series"
	cachedData, err := redis_lib.GetDataInRedis(c, cacheKey)

	if err == redis.Nil {
		log.Println("Cache miss for main page series. Fetching from TMDB...")

		apiKey := os.Getenv("TMDB_API")

		data, err := util.FetchTmdbExtraData(apiKey, url)
		if err != nil {
			log.Println("Error fetching series from TMDB:", err)
			return nil, err
		}

		fieldsToKeep := []string{"id", "poster_path", "title"}

		filteredSeries, err := util.FilterListData[interface{}](data, fieldsToKeep)

		jsonData, err := json.Marshal(filteredSeries)
		if err != nil {
			log.Println("Error marshalling data: ", err)
			return nil, err
		}

		err = redis_lib.SetDataInRedis(c, cacheKey, jsonData, 1*time.Hour)
		if err != nil {
			log.Println("Error caching series:", err)
			return nil, err
		}

		return filteredSeries, nil

	}

	var cachedSeries []map[string]interface{}
	err = json.Unmarshal([]byte(cachedData), &cachedSeries)
	if err != nil {
		log.Println("Error unmarshalling movies:", err)
	}

	return cachedSeries, nil
}
