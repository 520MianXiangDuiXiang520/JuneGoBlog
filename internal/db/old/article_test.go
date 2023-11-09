package old

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

func TestHasArticle(t *testing.T) {
	unExist := 10000
	exist := 29
	if HasArticle(unExist) {
		t.Error("Fail Test")
	}
	if !HasArticle(exist) {
		t.Error("Fail Test")
	}
}

func TestQueryArticleIDListFromCache(t *testing.T) {
	ids, err := queryArticleIDListFromCache(5, 11, 10)
	if err != nil {
		t.Error(err)
	}
	log.Println(ids)
}
