package main

import (
	"blog-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/posts", handlers.GetPosts)
	r.POST("/posts", handlers.CreatePost)
	r.Run()
}
