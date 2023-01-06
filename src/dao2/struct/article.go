package _struct

import "fmt"

type Article struct {
	ID         int     `json:"id" bson:"id"`
	Title      string  `json:"title" bson:"title"`
	Abstract   string  `json:"abstract" bson:"abstract"`
	Text       string  `json:"text" bson:"text"`
	CreateTime int64   `json:"createTime" bson:"create_time"`
	TagIds     []int64 `json:"tag_ids" bson:"tag_ids"`
}

func (a *Article) TName() string {
	return "article"
}

func (a *Article) String() string {
	return fmt.Sprintf("Article: %s(%d)", a.Title, a.ID)
}
