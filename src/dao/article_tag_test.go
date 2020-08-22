package dao

import (
	"log"
	"testing"
)

func TestQueryAllTagsByArticleID(t *testing.T) {
	tags := make([]Tag, 0)
	err := QueryAllTagsByArticleID(31, &tags)
	if err != nil {
		t.Error("QueryAllTagsByArticleID Error")
	}
	log.Println(len(tags))
	log.Println(tags[0].Name)
}

func TestQueryArticleTotalByTagID(t *testing.T) {
	total := QueryArticleTotalByTagID(16)
	if total != 1 {
		t.Error("")
	}
}

func TestQueryArticleTotalByTagIDFromCache(t *testing.T) {
	r, e := QueryArticleTotalByTagIDFromCache(15)
	if e != nil {
		t.Error(e)
	}
	log.Printf("r = [%v]", r)
	if r == 0 {
		t.Error("")
	}
}
