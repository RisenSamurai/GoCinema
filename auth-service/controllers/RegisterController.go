package controllers

import (
	"auth-service/models"
	"auth-service/mongo"
	"auth-service/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SignUp(c *gin.Context) {

	var user models.User

	log.Println("SignUp controller")

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user.Role = "user"

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user.Password = hashedPassword

	_, err = mongo.InsertItemInMongo(c, "Users", user)
	if err != nil {
		log.Println("Error inserting user", err)
	}

	c.JSON(200, gin.H{
		"code": 200,
	})

}
