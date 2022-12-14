package v1

import (
	"blog/internal/service"
	"github.com/gin-gonic/gin"
)

type Post struct {
	service service.Post
}

func NewPost() Post {
	postService := service.NewPost()
	return Post{postService}
}

func (p Post) Get(c *gin.Context) {
}

func (p Post) List(c *gin.Context) {
	var postList = p.service.List(c)
	c.JSON(200, postList)
}

func (p Post) ListByTag(c *gin.Context) {

}

func (p Post) ListByCategory(c *gin.Context) {

}
