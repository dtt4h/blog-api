package handlers

import (
	"blog-api/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{
		{
			ID:        1,
			Title:     "Первый пост про Golang",
			Content:   "Как устроен массив в Golang под капотом..",
			Published: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	c.JSON(http.StatusOK, posts)
}
