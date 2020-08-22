package dao

import (
	"JuneGoBlog/src/consts"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"testing"
)

func TestInitArticleIDListCache(t *testing.T) {
	_ = InitArticleIDListCache()
	rc := RedisPool.Get()
	defer rc.Close()
	articleList := make([]Article, 0)
	if err := QueryAllArticle(&articleList); err != nil {
		log.Println("InitArticleIDListCache Error!!")
		return
	}
	for index, article := range articleList {
		value, err := redis.String(rc.Do("LINDEX", consts.ArticleIDListCache, index))
		if err != nil {
			t.Error("Do Error")
		}
		if value != strconv.Itoa(article.ID) {
			t.Error("No Equal")
		}
	}
}

func TestInitTagArticleTotal(t *testing.T) {
	InitTagArticleTotal()
}

//func TestInitArticleInfoCache(t *testing.T) {
//	err := InitArticleInfoCache()
//	if err != nil {
//		t.Error(err)
//	}
//	rc := RedisPool.Get()
//	defer rc.Close()
//	articleList := make([]Article, 0)
//	if err := QueryAllArticle(&articleList); err != nil {
//		log.Println("InitArticleIDListCache Error!!")
//		return
//	}
//	for _, article := range articleList {
//		//"ID", "Title", "Abstract", "AuthorID", "CreateTime",
//		id, err := HGet(rc, consts.ArticleInfoHashCache + strconv.Itoa(article.ID), "ID")
//		DoError(err)
//		title, err := HGet(rc, consts.ArticleInfoHashCache + strconv.Itoa(article.ID), "Title")
//		DoError(err)
//		Abstract, err := HGet(rc, consts.ArticleInfoHashCache + strconv.Itoa(article.ID), "Abstract")
//		DoError(err)
//		AuthorID, err := HGet(rc, consts.ArticleInfoHashCache + strconv.Itoa(article.ID), "AuthorID")
//		DoError(err)
//		//CreateTime, err := HGet(rc, consts.ArticleInfoHashCache + strconv.Itoa(article.ID), "CreateTime")
//		if id != strconv.Itoa(article.ID) || title != article.Title || Abstract != article.Abstract ||
//					AuthorID != strconv.Itoa(article.AuthorID){
//			t.Error("No Equal!!!")
//
//		}
//	}
//
//}

func DoError(err error) {
	if err := QueryAllArticle(&articleList); err != nil {
		log.Println("InitArticleIDListCache Error!!")
		panic("END")
	}
}
