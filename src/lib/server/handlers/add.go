package handlers

import (
	"GoCinema/src/lib/server/database"
	"github.com/gin-gonic/gin"
	age2 "github.com/theTardigrade/golang-age"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"path/filepath"
	"time"
)

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
		actor.Images = append(actor.Images, filePath)
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
