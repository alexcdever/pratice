package model

type Tag struct {
	Model
	PostId int64
	Tag    string
}

func (Tag) TableName() string {
	return "tag"
}
