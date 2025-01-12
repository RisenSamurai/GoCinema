package controllers

import (
	"GoCinema/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAuthToken(c *gin.Context) {

}

func GetRegisterResponse(c *gin.Context) {
	data, err := services.SendRegisterData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, data)
}

func SignUpDataTest(c *gin.Context) {
	var data AuthData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Email: %s, Password: %s", data.Email, data.Password)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
