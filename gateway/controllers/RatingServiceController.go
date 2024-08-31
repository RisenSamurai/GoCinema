package controllers

import (
	"GoCinema/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainPageItems(c *gin.Context) {
	data, err := services.FetchTmdbMainPage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetTmdbPageItem(c *gin.Context) {
	data, err := services.FetchTmdbPageItem(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
