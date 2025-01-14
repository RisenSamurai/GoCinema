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

	var formData LoginForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Println(err.Error())
	}

	requestData, err := json.Marshal(formData)
	if err != nil {
		log.Println(err.Error())
	}

	resp, err := http.Post("http://127.0.0.1:8000/login", "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		log.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)

	return result, nil
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
