package routes

import (
	"GoCinema/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	crudService := router.Group("/edit")
	{
		crudService.GET("/push-actor")
	}

	ratingService := router.Group("/tmdb")
	{
		ratingService.GET("/items", controllers.GetMainPageItems)
		ratingService.GET("/page", controllers.GetTmdbPageItem)
	}
}
