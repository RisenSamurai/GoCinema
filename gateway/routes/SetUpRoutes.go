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

	regAuthService := router.Group("/auth")
	{
		regAuthService.POST("/login")
		regAuthService.POST("/register", controllers.GetRegisterResponse)
	}

	ratingService := router.Group("/tmdb")
	//ratingService.Use(middleware.AuthMiddleware())
	{
		ratingService.GET("/movie/:id", controllers.GetMoviePage)
		ratingService.GET("/series/:id", controllers.GetSeriesPage)
		ratingService.GET("/article/:id", controllers.GetArticlePage)
		ratingService.GET("/main-page-movies", controllers.GetMainPageMovies)
	}
}
