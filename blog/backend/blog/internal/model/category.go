package model

type Category struct {
	Model
	PostId   int64
	Category string
}

func (Category) TableName() string {
	return "category"
}
