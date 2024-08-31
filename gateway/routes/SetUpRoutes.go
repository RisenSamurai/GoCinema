package routes

import (
	"GoCinema/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	crudService := router.Group("/edit")
	{

	}

	ratingService := router.Group("/tmdb")
	{
		ratingService.GET("/", controllers.GetMainPageItems)
	}
}
