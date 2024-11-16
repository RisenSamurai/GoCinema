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
func GetMainPageMovies(c *gin.Context) {
	data, err := services.FetchMainPageMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}
