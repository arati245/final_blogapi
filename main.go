package main

import (
	"blogapp/database"
	"blogapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/posts", handlers.CreatePost)
	r.GET("/posts", handlers.GetPosts)
	r.GET("/posts/:id", handlers.GetPostByID)
	r.PUT("/posts/:id", handlers.UpdatePost)
	r.DELETE("/posts/:id", handlers.DeletePost)

	r.Run(":8080") // default port 8080
}
