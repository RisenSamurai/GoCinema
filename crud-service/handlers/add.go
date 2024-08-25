package handlers

import (
	"GoCinema/database"

	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	age2 "github.com/theTardigrade/golang-age"
)

func AddArticle(c *gin.Context) {
	var article database.Article
	log.Println("Inside AddArticle")

	const maxFileSize = 10 << 20

	err := c.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Printf("Encountered error during parsing form %s", err)
		return
	}

	file, err := c.FormFile("preview")
	if err != nil {
		fmt.Printf("Encountered error during getting file %s", err)
		return
	}

	article.Author = "Tom Cruise"
	article.Content = c.PostForm("content")
	article.Title = c.PostForm("title")
	article.Created = time.Now()

	if file != nil {

		uploadDir := filepath.Join("static", "images", "articles", article.Title)

		_, err := os.Stat(uploadDir)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(uploadDir, 0755)
				if err != nil {
					fmt.Printf("Encountered error during creating dir %s", err)
					return
				}
			}
		}

		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/webp": true,
		}

		if !allowedTypes[file.Header.Get("Content-Type")] {
			fmt.Printf("Image type is not allowed")
			return
		}

		ext := filepath.Ext(file.Filename)
		uniqueFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

		filePath := filepath.Join(uploadDir, uniqueFilename)

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			fmt.Printf("Encountered error saving file %s", err)
			return
		}

	}

	_, err = pushToDatabase(c, article, "Articles")
	{
		if err != nil {
			fmt.Printf("Encountered error during pushing to database: %s", err)
		}
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})

}

func AddActor(c *gin.Context) {
	log.Println("Got into AddActor!")

	var actor database.Actor

	const maxFileSize = 10 << 20

	err := c.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		log.Println("Error parsing multipart form")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["images"]

	log.Println(files)

	actor.Name = c.PostForm("name")
	actor.Surname = c.PostForm("lastName")

	birthdayStr := c.PostForm("birthday")

	birthday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		log.Println("Error parsing birthday")
		c.JSON(400, gin.H{
			"message": "Invalid date format" + err.Error(),
		})

		return
	}

	age := age2.CalculateToNow(birthday)
	actor.Age = age
	actor.Birthday = birthday
	actor.Gender = c.PostForm("gender")
	actor.Birthplace = c.PostForm("pob")
	actor.Bio = c.PostForm("biog")
	actor.Created = time.Now()

	uploadDir := filepath.Join("static", "images", "actors", actor.Name+actor.Surname+"/")

	err = os.Mkdir(uploadDir, 0755)
	if err != nil {
		log.Println("Error creating upload dir", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if files != nil {

		for _, file := range files {
			if !allowedTypes[file.Header.Get("Content-Type")] {
				c.JSON(400, gin.H{
					"message": "File type not allowed",
				})

				return
			}
			ext := filepath.Ext(file.Filename)
			uniqueFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

			filePath := filepath.Join(uploadDir, uniqueFilename)

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				log.Println("Error saving uploaded file", err)
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
				return
			}
			actor.Images = append(actor.Images, filePath)
		}
	}

	message, err := pushToDatabase(c, actor, "Actors")
	if err != nil {
		log.Println("Error pushing actor")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": message,
	})

}

func AddMovie(c *gin.Context) {

	var movie database.Movie

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		log.Println("Error parsing multipart form", err)
		c.JSON(400, gin.H{"error": "File too large"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing multipart form", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	movieDurationStr := c.PostForm("duration")
	movieDateStr := c.PostForm("releaseDate")
	movieBudgetStr := c.PostForm("budget")

	movieDuration, err := strconv.ParseFloat(movieDurationStr, 64)
	if err != nil {
		log.Println("Error parsing duration", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	movieBudget, err := strconv.ParseFloat(movieBudgetStr, 64)
	if err != nil {
		log.Println("Error parsing budget", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	movieDate, err := time.Parse("2006-01-02", movieDateStr)
	if err != nil {
		log.Println("Error parsing date", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	movie.Name = c.PostForm("name")
	movie.ReleaseDate = movieDate
	movie.Duration = movieDuration
	movie.Budget = movieBudget
	movie.Description = c.PostForm("description")
	movie.Year = c.PostForm("year")
	movie.Directors = c.PostFormArray("directors")
	movie.Writers = c.PostFormArray("writers")
	movie.Producers = c.PostFormArray("producers")
	movie.Actors = c.PostFormArray("actors")
	movie.Countries = c.PostFormArray("countries")
	movie.Genres = c.PostFormArray("genres")
	movie.Cameras = c.PostFormArray("cameras")
	movie.Editors = c.PostFormArray("editors")
	movie.Keywords = c.PostFormArray("keywords")
	movie.Language = c.PostForm("language")

	uploadDir := filepath.Join("images", "movie", movie.Name)

	err = DirExists(uploadDir)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(400, gin.H{
			"message": "No files uploaded",
		})
		return
	}

	if files != nil {

		for _, fileHeader := range files {
			if !allowedTypes[fileHeader.Header.Get("Content-Type")] {
				c.JSON(400, gin.H{
					"message": "File content type not allowed",
				})
				return
			}

			ext := filepath.Ext(fileHeader.Filename)
			uniqueFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

			filePath := filepath.Join(uploadDir, uniqueFilename)
			if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
				log.Println("Error saving file:", err)
				c.JSON(500, gin.H{
					"message": "Failed to save file",
				})
				return
			}

			movie.Images = append(movie.Images, filePath)
		}
	}

	file, err := c.FormFile("poster")
	if err != nil {
		log.Println("Error retrieving poster", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}
	if file != nil {

		if !allowedTypes[file.Header.Get("Content-Type")] {
			c.JSON(400, gin.H{
				"message": "File content type not allowed",
			})
			return
		}
		// Create the directory for posters if it doesn't exist
		posterUploadDir := filepath.Join("static", "images", "posters", movie.Name)
		if err := os.MkdirAll(posterUploadDir, 0755); err != nil {
			log.Println("Error creating poster directory:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poster directory"})
			return
		}

		// Generate a unique filename
		ext := filepath.Ext(file.Filename)
		uniqueFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

		// Create the full file path
		posterFilePath := filepath.Join(posterUploadDir, uniqueFilename)

		// Save the file
		if err := c.SaveUploadedFile(file, posterFilePath); err != nil {
			log.Println("Error saving poster file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save poster file"})
			return
		}

		movie.Poster = filepath.Join("images", "posters", movie.Name, uniqueFilename)
	}

	_, err = pushToDatabase(c, movie, "Movies")
	if err != nil {
		log.Println("Error pushing movie")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Movie successfully uploaded",
	})

}

func pushMovie(c *gin.Context, movie database.Movie) (string, error) {

	client, err := database.Cn()
	if err != nil {
		log.Println("Error getting cn", err)
	}

	collection := client.Database("GoCinema").Collection("Movies")

	bsonMovie := bson.M{
		"name":        movie.Name,
		"year":        movie.Year,
		"duration":    movie.Duration,
		"releaseDate": movie.ReleaseDate,
		"description": movie.Description,
		"writers":     movie.Writers,
		"directors":   movie.Directors,
		"producers":   movie.Producers,
		"editors":     movie.Editors,
		"cameras":     movie.Cameras,
		"genres":      movie.Genres,
		"actors":      movie.Actors,
		"budget":      movie.Budget,
		"language":    movie.Language,
		"countries":   movie.Countries,
		"keywords":    movie.Keywords,
		"images":      movie.Images,
		"poster":      movie.Poster,
		"created":     time.Now(),
	}

	_, err = collection.InsertOne(c.Request.Context(), bsonMovie)
	if err != nil {
		log.Println("Error inserting movie", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return "", err
	}

	return "Movie added", nil
}

func pushToDatabase[T any](c *gin.Context, obj T, collectionName string) (string, error) {
	client, err := database.Cn()
	if err != nil {
		fmt.Printf("Error getting connection")
		return "Error getting connection", err
	}

	collection := client.Database("GoCinema").Collection(collectionName)

	_, err = collection.InsertOne(c.Request.Context(), obj)
	if err != nil {
		fmt.Println("Error inserting object", err)
		return "Encountered error: ", err
	}

	return "Object inserted!", nil

}
