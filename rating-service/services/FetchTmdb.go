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

//Deprecated code

/*

func FetchRating(c *gin.Context, tmbdbID string) (interface{}, error) {

	log.Println("Inside FetchRating")
	id := c.Param("id")
	cacheKey := "movie:" + id

	if redisClient == nil {
		log.Println("Redis client is nil. Initializing...")
		initRedis()
	}

	ctx := c.Request.Context()
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Data found in cache")
		var movieData map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &movieData); err != nil {
			log.Printf("Failed to parse cached data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse cached data"})
			return "", err
		}
		return movieData, err
	} else if !errors.Is(err, redis_lib.Nil) {
		log.Printf("Redis error: %v", err)
	} else {
		log.Println("Data not found in cache")
	}

	apiKey := os.Getenv("TMDB_API")
	if apiKey == "" {
		log.Println("API key not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not found"})
		return "", err
	}

	// Data not in cache or error occurred, fetch from API
	movieData, err := util.FetchTmdbExtraData(c, apiKey, tmbdbID)
	if err != nil {
		log.Printf("Failed to fetch movie data: %v", err)
		return "", err
	}

	// Process the data if needed
	processedData := gin.H{
		"id":           movieData["id"],
		"title":        movieData["title"],
		"rating":       movieData["vote_average"],
		"vote_count":   movieData["vote_count"],
		"overview":     movieData["overview"],
		"popularity":   movieData["popularity"],
		"runtime":      movieData["runtime"],
		"release_date": movieData["release_date"],
		"revenue":      movieData["revenue"],
	}

	// Cache the processed data in Redis
	jsonData, err := json.Marshal(processedData)
	if err != nil {
		log.Printf("Failed to marshal data for caching: %v", err)
	} else {
		err = setDataInRedis(c, cacheKey, jsonData, 1*time.Hour) // Cache for 1 hour
		if err != nil {
			log.Printf("Redis caching error: %v", err)
		} else {
			log.Println("Successfully cached data in Redis")
		}
	}

	return processedData, nil

}




func FetchMovie(c *gin.Context) (interface{}, error) {
	movieID := c.Param("id")

	movie, err := database.FetchItemFromMongo(c, movieID)
	if err != nil {
		return "Failed to fetch item from Mongo", err
	}

	ratings, err := FetchRating(c, movie.TmdbId)
	if err != nil {
		return "Failed to fetch rating", err
	}

	items := gin.H{
		"movie":   movie,
		"ratings": ratings,
	}

	log.Println("Items to Gateway: ", items)

	return items, nil

}

*/

/*

func FetchItems(c *gin.Context) (interface{}, error) {
	var movies []database.Movie
	var articles []database.Article

	// Fetch movies from MongoDB
	movies, err := database.FetchAnyFromMongo[database.Movie](c.Request.Context(), "Movies")
	if err != nil {
		log.Println("Error fetching movies from Mongo:", err)
		return nil, fmt.Errorf("error fetching movies from Mongo")
	}

	// Fetch articles from MongoDB
	articles, err = database.FetchAnyFromMongo[database.Article](c.Request.Context(), "Articles")
	if err != nil {
		log.Println("Error fetching articles from Mongo:", err)
		return nil, fmt.Errorf("error fetching articles from Mongo")
	}

	// Log successful fetch
	log.Println("Successfully fetched movies and articles from Mongo")

	items := gin.H{
		"movies":   movies,
		"articles": articles,
	}

	return items, nil
}



*/

func FetchMoviePage(c *gin.Context) (database.DetailedMovie, error) {
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

		var detailedMovie database.DetailedMovie

		genres, err := util.ParseGenres[util.Genre](data["genres"])
		if err != nil {
			log.Println("FetchMoviePage: error parsing genres: ", err)
		}
		countries, err := util.ParseGenres[any](data["production_countries"])
		if err != nil {
			log.Println("FetchMoviePage: error parsing genres: ", err)
		}
		productionCompanies, err := util.ParseGenres[any](data["production_companies"])
		if err != nil {
			log.Println("FetchMoviePage: error parsing genres: ", err)
		}

		detailedMovie = database.DetailedMovie{
			Id:                  int(data["id"].(float64)),
			Title:               data["title"].(string),
			Budget:              data["budget"].(float64),
			Language:            data["original_language"].(string),
			Genres:              genres,
			Country:             countries,
			ProductionCompanies: productionCompanies,
			ReleaseDate:         data["release_date"].(string),
			Duration:            data["runtime"].(float64),
			Description:         data["overview"].(string),
			Popularity:          data["popularity"].(float64),
			VoteAverage:         data["vote_average"].(float64),
			VoteCount:           data["vote_count"].(float64),
			Revenue:             data["revenue"].(float64),
			Status:              data["status"].(string),
			Tagline:             data["tagline"].(string),
			Poster:              data["poster_path"].(string),
		}

		jsonData, err := json.Marshal(detailedMovie)
		if err != nil {
			log.Println("FetchMoviePage: error marshalling detailed movie: ", err)
		}

		err = redis_lib.SetDataInRedis(c, cacheKey, jsonData, 1*time.Hour)
		if err != nil {
			log.Println("FetchMoviePage: error caching movie details: ", err)
		}

		return detailedMovie, nil

	}

	var cachedMovie database.DetailedMovie

	err = json.Unmarshal([]byte(cachedData), &cachedMovie)

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
