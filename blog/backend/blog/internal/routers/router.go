package routers

import (
	v1 "blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	post := v1.NewPost()
	tag := v1.NewTag()
	category := v1.NewCategory()
	apiV1 := r.Group("/api/v1")

	// get all posts
	postsApi := apiV1.Group("/posts")
	postsApi.GET("", post.List)
	categoriesApi := postsApi.Group("/categories")
	categoriesApi.GET("", category.Get)
	tagsApi := postsApi.Group("/tags")
	tagsApi.GET("", tag.Get)
	// other info
	{
		// get the special post
		postsApi.GET("/:id", post.Get)
		// get the posts by tag
		tagsApi.GET("/:tag", post.ListByTag)
		// get the posts by category
		categoriesApi.GET("/:category", post.ListByCategory)
	}

	return r
}
