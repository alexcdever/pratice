package v1

import "github.com/gin-gonic/gin"

type Category struct {
}

func NewCategory() Category {
	return Category{}
}

func (c Category) Get(g *gin.Context) {

}
func (c Category) List(g *gin.Context) {

}
