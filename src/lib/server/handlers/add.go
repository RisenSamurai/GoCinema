package handlers

import (
	"GoCinema/src/lib/server"
	"GoCinema/src/lib/server/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	age2 "github.com/theTardigrade/golang-age"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieCollection interface {
	InsertOne(ctx context.Context, movie interface{}) (*mongo.InsertOneResult, error)
}

type Handler struct {
	Client *mongo.Client
}

func NewHandler(client *mongo.Client) *Handler {
	return &Handler{Client: client}
}

func (h *Handler) AddActor(c *gin.Context) {
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

	log.Println("Birthday: ", birthdayStr)
	log.Println("Name:", actor.Name)

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

	dir := filepath.Dir("./static/images/actors/")
	uploadDir := filepath.Join(dir, actor.Name+actor.Surname+"/")

	err = os.Mkdir(uploadDir, 0755)
	if err != nil {
		log.Println("Error creating upload dir", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, file := range files {
		filePath := filepath.Join(uploadDir, file.Filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Println("Error saving uploaded file", err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		//actor.Images = append(actor.Images, filePath)
	}

	message, err := h.pushActor(c, actor)
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

func (h *Handler) AddMovie(c *gin.Context) {

	var movie database.Movie

	log.Println("-We are inside AddMovie")

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

	movieDuration, err := strconv.ParseFloat(movieDurationStr, 64)
	if err != nil {
		log.Println("Error parsing duration", err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
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
	movie.Description = c.PostForm("description")
	movie.Year = c.PostForm("year")

	dir := filepath.Dir("./static/images/movies/")
	uploadDir := filepath.Join(dir, movie.Name+"/")

	err = server.DirExists(uploadDir)
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

	for _, fileHeader := range files {
		if !allowedTypes[fileHeader.Header.Get("Content-Type")] {
			c.JSON(400, gin.H{
				"message": "File content type not allowed",
			})
			return
		}

		filePath := filepath.Join(uploadDir, fileHeader.Filename)
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			log.Println("Error saving file:", err)
			c.JSON(500, gin.H{
				"message": "Failed to save file",
			})
			return
		}

		movie.Images = append(movie.Images, filePath)
	}

	file, err := c.FormFile("poster")
	if err != nil {
		log.Println("Error retrieving poster")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	if len(files) == 0 {
		c.JSON(400, gin.H{
			"message": "No files uploaded",
		})
		return
	}

	posterUploadDir := filepath.Join(".static/images/posters/", movie.Name)
	posterFilePath := filepath.Join(posterUploadDir, file.Filename)

	err = server.DirExists(posterFilePath)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	movie.Poster = posterFilePath
	if err := c.SaveUploadedFile(file, movie.Poster); err != nil {
		log.Println("Error saving file:", err)
	}

	_, err = h.pushMovie(c, movie)
	if err != nil {
		log.Println("Error pushing movie")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Movie" + movie.Name + "successfully uploaded",
	})

}

func (h *Handler) pushMovie(c *gin.Context, movie database.Movie) (string, error) {

	collection := h.Client.Database("GoCinema").Collection("Movies")

	bsonMovie := bson.M{
		"name":        movie.Name,
		"year":        movie.Year,
		"duration":    movie.Duration,
		"releaseDate": movie.ReleaseDate,
		"description": movie.Description,
		"images":      movie.Images,
		"poster":      movie.Poster,
		"created":     time.Now(),
	}

	_, err := collection.InsertOne(c.Request.Context(), bsonMovie)
	if err != nil {
		log.Println("Error inserting movie", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return "", err
	}

	return "Movie added", nil
}

func (h *Handler) pushActor(c *gin.Context, actor database.Actor) (string, error) {

	collection := h.Client.Database("GoCinema").Collection("Actors")

	log.Println("got into push Actor!")
	_, err := collection.InsertOne(c.Request.Context(), actor)
	if err != nil {
		log.Println("Error pushing actor", err)
		c.JSON(500, gin.H{

			"message": err.Error(),
		})
	}

	return "Actor added", nil

}
