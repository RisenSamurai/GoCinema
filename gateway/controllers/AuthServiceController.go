package controllers

import (
	"GoCinema/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuthToken(c *gin.Context) {

	data, err := services.SendLoginData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

func GetRegisterResponse(c *gin.Context) {
	data, err := services.SendRegisterData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, data)
}
