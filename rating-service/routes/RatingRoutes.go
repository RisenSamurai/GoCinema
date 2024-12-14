package routes

import (
	"github.com/gin-gonic/gin"
	"rating_microservice/controllers"
)

func SetupRatingRoutes(router *gin.Engine) {
	tmdb := router.Group("/fetch")
	{
		tmdb.GET("/main-page-movies", controllers.GetMainPageMovies)
		tmdb.GET("/main-page-series", controllers.GetMainPageSeries)
		tmdb.GET("/series/:id", controllers.GetSeriesPage)
		tmdb.GET("/articles", controllers.GetMainPageArticles)
		tmdb.GET("/article/:id")
		tmdb.GET("/movie/:id", controllers.GetMoviePage)
	}
}
