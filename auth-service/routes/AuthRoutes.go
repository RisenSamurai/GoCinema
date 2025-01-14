package routes

import (
	"auth-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
	}
}
