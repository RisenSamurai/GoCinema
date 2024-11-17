package controllers

import (
	"GoCinema/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetMoviePage(c *gin.Context) {
	data, err := services.FetchPageMovie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	log.Println("Gateway sends MoviePage data: ", data)

	c.JSON(http.StatusOK, data)
}

func GetMainPageMovies(c *gin.Context) {
	data, err := services.FetchMainPageMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": data,
	})
}
