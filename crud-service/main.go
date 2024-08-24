package main

import (
	"GoCinema/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/add-actor", handlers.AddActor)
	r.POST("/add-movie", handlers.AddMovie)
	r.POST("/add-article", handlers.AddArticle)
	r.GET("/fetch-main-page-items", handlers.GetItems)
	r.GET("/fetch-movie/:id", handlers.GetMovie)

	r.Run("localhost:8082")

}
