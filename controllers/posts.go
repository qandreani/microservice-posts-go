package controllers

import (
	"API-Webservice/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePostInput struct {
	Text string `json:"post" binding:"required"`
}

type UpdatePostInput struct {
	Text string `json:"post"`
}

// GET /posts
// Find all posts
func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GET /posts/:id
// Find a post
func FindPost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// POST /posts
// Create new post
func CreatePost(c *gin.Context) {
	// Validate input
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	post := models.Post{Text: input.Text}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /posts/:id
// Update a post
func UpdatePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&post).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
