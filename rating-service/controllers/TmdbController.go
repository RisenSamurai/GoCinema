package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rating_microservice/services"
)

/*
	func GetMainPageItems(c *gin.Context) {
		data, err := services.FetchItems(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, data)
	}
*/

func GetMoviePage(c *gin.Context) {
	data, err := services.FetchMoviePage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, data)
}

func GetMainPageMovies(c *gin.Context) {
	data, err := services.FetchMainPageMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}
