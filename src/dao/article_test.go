package dao

import (
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

func TestGetArticleIDListCacheIndexByID(t *testing.T) {
	_, _ = getArticleIDListCacheIndexByID(1)
}
