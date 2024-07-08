package server

import (
	"log"
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

func WriteImages(dir string, images map[string]string) error {

	if err := DirExists(dir); err != nil {
		return err
	}

	for filename, content := range images {
		filePath := filepath.Join(dir, filename)
		if err := WriteImage(filePath, filename, content); err != nil {
			return err
		}
	}

	return nil

}

func WriteImage(dir string, fileName, content string) error {
	if err := DirExists(dir); err != nil {
		return err
	}

	filePath := filepath.Join(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating file: ", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("Error writing to file: ", err)
	}

	log.Println("Created file: ", filePath)

	return nil
}
