package service

import (
	"github.com/gin-gonic/gin"
	"modern-blog/internal/dao"
	"modern-blog/internal/model"
)

type Post struct {
}

func (p Post) List(c *gin.Context) (result []model.Post) {
	dao.DbConnection.Select("id", "title", "draft", "filename", "created_at", "updated_at").Find(&result)
	return result
}

func NewPost() Post {
	return Post{}
}
