package handlers

import (
	"GoCinema/src/lib/server"
	"GoCinema/src/lib/server/database"
	"context"
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
		c.JSON(500, gin.H{
			"message": err.Error(),
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
	const maxFileSize = 10 << 20

	err := c.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		log.Println("Error parsing multipart form")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	releaseDateStr := c.PostForm("releaseDate")
	durationStr := c.PostForm("duration")

	duration, err := strconv.ParseFloat(durationStr, 32)
	if err != nil {
		log.Println("Error parsing duration")
	}

	releaseDate, err := time.Parse("2006-01-02", releaseDateStr)
	if err != nil {
		log.Println("Error parsing birthday")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	movie.Name = c.PostForm("movie-name")
	movie.Year = c.PostForm("year")
	movie.Directors = c.PostFormArray("director")
	movie.Writers = c.PostFormArray("writers")
	movie.Producers = c.PostFormArray("producers")
	movie.Editors = c.PostFormArray("editors")
	movie.Cameras = c.PostFormArray("cameras")
	movie.Genres = c.PostFormArray("genres")
	movie.ReleaseDate = releaseDate
	movie.Countries = c.PostFormArray("countries")
	movie.Duration = duration
	movie.Description = c.PostForm("description")
	movie.Actors = c.PostFormArray("actors")
	movie.Created = time.Now()

	dir := filepath.Dir("./static/images/movies/")
	uploadDir := filepath.Join(dir, movie.Name+"/")

	err = server.DirExists(uploadDir)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error retrieving multipart form")
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

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
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

	movie.Poster = filepath.Join(uploadDir, file.Filename)

	_, err = h.pushMovie(c, movie)
	if err != nil {
		log.Println("Error pushing movie")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Movie added",
	})

}

func (h *Handler) pushMovie(c *gin.Context, movie database.Movie) (string, error) {

	collection := h.Client.Database("GoCinema").Collection("Movies")

	_, err := collection.InsertOne(c.Request.Context(), movie)
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
