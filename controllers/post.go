package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nadduli/Trove/dto"
	"github.com/nadduli/Trove/initializers"
	"github.com/nadduli/Trove/models"
)

func CreatePost(c *gin.Context) {
	var body dto.CreatePostDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}
	if result := initializers.DB.Create(&post); result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Post created successfully",
		"data":    post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"message": "retrieved all posts successfully",
		"posts":   posts,
	})
}
