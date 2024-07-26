package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func DirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {

			log.Println("Error creating directory: ", err)
			return err
		}

		log.Println("Created directory: ", path)
	} else if err != nil {
		log.Println("Error checking directory: ", err)
	}

	return nil
}

func UploadImage(c *gin.Context, dir string, postValue string) (string, error) {

	var file *multipart.FileHeader

	file, err := c.FormFile(postValue)
	if err != nil {
		log.Println("Error getting file: ", err)
		return "", err
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedTypes[file.Header.Get("Content-Type")] {
		c.JSON(500, gin.H{
			"message": "File content type not allowed",
		})

		return "", err
	}

	err = DirExists(dir)
	if err != nil {
		log.Println("Error checking directory: ", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return "", err
	}

	filename := filepath.Join(dir, file.Filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return "", err
	}

	return filename, nil
}

func UploadImages(c *gin.Context, dir string, postValues []string) ([]string, error) {
	var paths []string
	for _, postValue := range postValues {
		path, err := UploadImage(c, dir, postValue)
		if err != nil {
			return paths, err
		}
		paths = append(paths, path)

		// Save the image path to the database
		err = saveImagePathToDB(path)
		if err != nil {
			log.Println("Error saving image path to database: ", err)
			return paths, err
		}
	}
	return paths, nil
}

func saveImagePathToDB(path string) error {
	// Implement your database logic here
	log.Println("Saving image path to database: ", path)
	return nil
}
