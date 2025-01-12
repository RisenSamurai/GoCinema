package services

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func SendLoginData(c *gin.Context) (interface{}, error) {

	var obj interface{}

	return obj, nil
}

func SendRegisterData(c *gin.Context) (interface{}, error) {

	var formData LoginForm

	log.Println("Data received from Front End")

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	requestData, err := json.Marshal(formData)
	if err != nil {
		return nil, err
	}
	log.Println("Marshaling..")

	resp, err := http.Post("http://localhost:8082/auth/sign-up",
		"application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
