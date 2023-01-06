package _struct

import "fmt"

type Tag struct {
	ID         int64  `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
	Total      int64  `json:"article_total" bson:"total"`
}

func (t *Tag) TName() string {
	return "tag"
}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag: %s(%d)-%d", t.Name, t.ID, t.Total)
}
