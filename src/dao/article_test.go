package dao

import (
	"log"
	"testing"
)

func TestQueryArticleIDListFromCache(t *testing.T) {
	_, _ = queryArticleIDListFromCache(1, 4)
}

func TestQueryArticleInfoListFromCache(t *testing.T) {
	ids := []int{
		99, 199, 299,
	}
	res, _ := getArticleInfoListFromCache(ids)
	log.Println(res)
}

func TestQueryArticleInfoByID(t *testing.T) {
	article := Article{}
	_ = QueryArticleInfoByID(10, &article)
	log.Printf(article.Title)
}
