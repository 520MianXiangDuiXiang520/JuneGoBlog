package dao

import (
	"log"
	"testing"
)

func TestQueryAllArticleByTagID(t *testing.T) {
	r, _ := QueryAllArticleByTagID(15)
	log.Println(r)
}
