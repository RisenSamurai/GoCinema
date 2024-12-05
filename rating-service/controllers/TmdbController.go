package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rating_microservice/services"
)

func GetMoviePage(c *gin.Context) {
	data, err := services.FetchMoviePage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, data)
}

func GetMainPageMovies(c *gin.Context) {
	data, err := services.FetchMainPageMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}

func GetMainPageSeries(c *gin.Context) {
	data, err := services.FetchMainPageSeries(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}
