package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rating_microservice/services"
)

func GetRating(c *gin.Context) {
	data, err := services.FetchRating(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
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
