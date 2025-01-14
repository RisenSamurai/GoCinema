package controllers

import (
	"auth-service/models"
	"auth-service/mongo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func SignIn(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbUser, err := mongo.GetUserFromMongo(c, "Users", user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if dbUser.Email == user.Email {

	}

}

func GenerateToken(c *gin.Context) {
	secretKey := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"user_id": 1,
		"role":    "user",
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})

}
