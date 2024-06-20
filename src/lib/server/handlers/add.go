package handlers

import (
	"GoCinema/src/lib/server/database"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

func AddActor(c *gin.Context) {
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

}

func pushActor(actor database.Actor) {

}
