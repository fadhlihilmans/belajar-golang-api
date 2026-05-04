package main

import (
	"belajar-golang-api/controllers"
	"belajar-golang-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()


	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// CATEGORIES
	router.GET("/api/categories", controllers.FindCategories)
	router.POST("/api/categories", controllers.StoreCategory)
	router.GET("/api/categories/:id", controllers.FindCategoryById)
	router.PUT("/api/categories/:id", controllers.UpdateCategory)
	router.DELETE("/api/categories/:id", controllers.DeleteCategory)

	// POSTS
	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/posts", controllers.StorePost)
	router.GET("/api/posts/:id", controllers.FindPostById)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)


	router.Run(":3000")
}
