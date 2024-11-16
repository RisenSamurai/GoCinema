package routes

import (
	"github.com/gin-gonic/gin"
	"rating_microservice/controllers"
)

func SetupRatingRoutes(router *gin.Engine) {
	tmdb := router.Group("/fetch")
	{
		tmdb.GET("/main-page-movies", controllers.GetMainPageMovies)
		tmdb.GET("/movie/:id")
	}
}
