package main

import (
	"simple-go-service/controllers"
	"simple-go-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	v1 := router.Group("/v1")

	v1.GET("/movies", controllers.FindMovies)
	v1.POST("movies", controllers.CreateMovie)
	v1.GET("/movies/:id", controllers.FindMovie)
	v1.PATCH("/movies/:id", controllers.UpdateMovie)
	v1.DELETE("/movies/:id", controllers.DeleteMovie)

	err := router.Run(":5000")
	if err != nil {
		return
	}

}
