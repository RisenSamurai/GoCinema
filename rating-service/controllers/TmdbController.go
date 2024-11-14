package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rating_microservice/services"
)

func GetMovie(c *gin.Context) {
	data, err := services.FetchMovie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}

func GetMainPageItems(c *gin.Context) {
	data, err := services.FetchItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}
