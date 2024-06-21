package handlers

import (
	"GoCinema/src/lib/server/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	var actor database.Actor

	const maxFileSize = 10 << 20

	err := c.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["images"]

	actor.Name = c.PostForm("name")
	actor.Surname = c.PostForm("lastName")

	birthdayStr := c.PostForm("birthday")

	birthday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	actor.Birthday = birthday
	actor.Gender = c.PostForm("gender")
	actor.Birthplace = c.PostForm("pob")
	actor.Bio = c.PostForm("bio")

	uploadDir := "static/images/actors/"

	err = os.Mkdir(filepath.Join(uploadDir, actor.Name), 0755)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	for _, file := range files {
		filePath := filepath.Join(uploadDir, actor.Name, file.Filename)

		actor.Images = append(actor.Images, filePath)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

	}

	message, err := h.pushActor(c, actor)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": message,
		})
	}

}

func (h *Handler) pushActor(c *gin.Context, actor database.Actor) (string, error) {

	collection := h.Client.Database("GoCinema").Collection("actors")

	_, err := collection.InsertOne(c.Request.Context(), actor)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	return "Actor added", nil

}
