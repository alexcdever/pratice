package model

type Post struct {
	Model
	Title    string `json:"title"`
	Draft    bool   `json:"draft"`
	Content  string `json:"content"`
	Filename string
	Md5      string
}

func (Post) TableName() string {
	return "post"
}
