package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Blog represents a blog post with a title, description, content, author, status.
// swagger:model
type Blog struct {
	// The ID of the blog
	//
	// example: 1
	ID int `json:"id"`

	// The title of the blog
	// required: true
	// example: My first blog
	Title string `json:"title"`

	// The description of the blog
	//
	// example: This is my first blog
	Description string `json:"description"`

	// The content of the blog
	// required: true
	// example: This is the content of my first blog
	Content string `json:"content"`

	// The author of the blog
	// required: true
	// example: John Doe
	Author string `json:"author"`

	// The publication status of the blog
	// required: true
	// example: true
	IsPublished bool `json:"isPublished"`
}

var blogs = []Blog{
	{
		ID:          1,
		Title:       "My first blog",
		Description: "This is my first blog",
		Content:     "This is the Content of my first blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          2,
		Title:       "My second blog",
		Description: "This is my second blog",
		Content:     "This is the Content of my second blog",
		Author:      "Jane Doe",
		IsPublished: false,
	},
	{
		ID:          3,
		Title:       "My third blog",
		Description: "This is my third blog",
		Content:     "This is the Content of my third blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          4,
		Title:       "My fourth blog",
		Description: "This is my fourth blog",
		Content:     "This is the Content of my fourth blog",
		Author:      "Jane Doe",
		IsPublished: true,
	},
}

func (b *Blog) GetBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blogs, "message": "success"})
}

func (b *Blog) GetBlog(c *gin.Context) {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": err.Error()})
		return
	}

	for _, blog := range blogs {
		if blog.ID == intID {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blog, "message": "success"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": ""})
}

func (b *Blog) CreateBlog(c *gin.Context) {
	var blog Blog

	blog.ID = len(blogs) + 1
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": "", "message": err.Error()})
		return
	}

	blogs = append(blogs, blog)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blog, "message": "success"})
}

func (b *Blog) UpdateBlog(c *gin.Context) {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": err.Error()})
		return
	}

	for index, blog := range blogs {
		if blog.ID == intID {
			var update_blog Blog
			err := c.BindJSON(update_blog)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": err.Error()})

				return
			}

			update_blog.ID = intID
			blogs[index] = update_blog

			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": update_blog, "message": "success"})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": "not found"})
	}
}

func (b *Blog) DeleteBlog(c *gin.Context) {
	intID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": "", "message": err.Error()})
		return
	}

	for index, blog := range blogs {
		if blog.ID == intID {
			blogs = append(blogs[:index], blogs[index+1:]...)

			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "", "message": "success"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": "not found"})
}
