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
		ratingService.GET("/movie/:id", controllers.GetMoviePage)
		ratingService.GET("/series/:id", controllers.GetSeriesPage)
		ratingService.GET("/article/:id")
		ratingService.GET("/main-page-movies", controllers.GetMainPageMovies)
	}
}
