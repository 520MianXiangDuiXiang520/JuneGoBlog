package dao

import (
	"log"
	"testing"
	"time"
)

func TestUpdateArticle(t *testing.T) {
	_ = UpdateArticle(58, &Article{
		Title:      "1Update Test",
		CreateTime: time.Now(),
		Text:       "Test",
		AuthorID:   1,
	})
}

func TestQueryArticleInfoByID(t *testing.T) {
	a := Article{}
	err := QueryArticleInfoByID(48, &a)
	if err != nil {
		t.Error("查询失败！")
	}
	log.Println(a)
}
