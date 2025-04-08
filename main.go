package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nadduli/Trove/controllers"
	"github.com/nadduli/Trove/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/", healthCheck)
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.GetPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"message": "welcome to Blog API",
	})
}
