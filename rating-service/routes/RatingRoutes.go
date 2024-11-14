package routes

import (
	"github.com/gin-gonic/gin"
	"rating_microservice/controllers"
)

func SetupRatingRoutes(router *gin.Engine) {
	tmdb := router.Group("/fetch")
	{
		tmdb.GET("/movie/:id", controllers.GetMovie)
		tmdb.GET("/fetch-main-page-items", controllers.GetMainPageItems)
	}
}
