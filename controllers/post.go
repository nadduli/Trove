package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nadduli/Trove/dto"
	"github.com/nadduli/Trove/initializers"
	"github.com/nadduli/Trove/models"
	"gorm.io/gorm"
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

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, "id = ?", id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID format",
		})
		return
	}

	var body dto.UpdatePostDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	updates := models.Post{
		Title:   body.Title,
		Content: body.Content,
	}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&post).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update post",
		})
		return
	}

	initializers.DB.First(&post, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID format",
		})
		return
	}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		var post models.Post
		if err := tx.First(&post, "id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&post).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to delete post",
				"details": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
