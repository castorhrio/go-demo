package main

import (
	"blog-api/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	blog := &routes.Blog{}

	route := gin.Default()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	route.Use(corsConfig)
	route.GET("/blogs", blog.GetBlogs)
	route.GET("/blogs/:id", blog.GetBlog)
	route.POST("/blogs", blog.CreateBlog)
	route.PUT("/blogs/:id", blog.UpdateBlog)
	route.DELETE("/blogs/:id", blog.DeleteBlog)

	route.Run("localhost:8057")
}
