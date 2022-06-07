package main

import (
	"API-Webservice/controllers"
	"API-Webservice/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/posts", controllers.FindPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.POST("/posts", controllers.CreatePost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Run the server
	r.Run()
}
