package dao

import (
	"JuneGoBlog/src/consts"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
	"strconv"
)

// 用来预热缓存
var articleList []Article

func init() {
	articleList = make([]Article, 0)
	if err := QueryAllArticle(&articleList); err != nil {
		log.Println("InitArticleIDListCache Error!!")
		return
	}
}

// 文章列表预热
//   1. 预热 文章ID 列表
//   2. 预热 文章简单信息

func InitArticleIDListCache() error {
	// 1. 更新缓存中的 articleIDList
	rc := RedisPool.Get()
	defer rc.Close()
	for index, article := range articleList {
		value, err := redis.String(rc.Do("LINDEX", consts.ArticleIDListCache, index))
		if err != nil {
			log.Println("Get Error!")
			return err
		}
		if value == "" {
			err := rc.Send("RPUSH", consts.ArticleIDListCache, article.ID)
			if err != nil {
				log.Println("插入ID失败", err)
				return err
			}
		} else if value != strconv.Itoa(article.ID) {
			if err := rc.Send("lset", consts.ArticleIDListCache,
				int64(index), article.ID); err != nil {
				log.Println("LSET 执行失败")
			}
		}
	}
	return nil
}

func HGet(conn redis.Conn, key, field string) (string, error) {
	v, err := redis.String(conn.Do("HGET", key, field))
	return v, err
}

func InitArticleInfoCache() error {
	rc := RedisPool.Get()
	defer rc.Close()
	values := [5]string{
		"ID", "Title", "Abstract", "AuthorID", "CreateTime",
	}
	for _, article := range articleList {
		for _, value := range values {
			v, err := HGet(rc, consts.ArticleInfoHashCache+strconv.Itoa(article.ID), value)
			//if err != nil {
			//	log.Println("HGET Error!", err)
			//	return err
			//}
			if v == "" || err == redis.ErrNil {
				immutable := reflect.ValueOf(article)
				val := immutable.FieldByName(value)
				_, err := rc.Do("HSET",
					consts.ArticleInfoHashCache+strconv.Itoa(article.ID),
					value, val)
				if err != nil {
					log.Println("HSET Error!")
					return err
				}
			}
		}

	}
	return nil
}

func InitTagArticleTotal() {
	rc := RedisPool.Get()
	defer rc.Close()

	allTags := make([]Tag, 0)
	_ = QueryAllTagsOrderByTime(&allTags)

	for _, tag := range allTags {
		total, _ := QueryArticleTotalByTagID(tag.ID)
		rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "ID", tag.ID)
		rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "Name", tag.Name)
		rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "CreateTime", tag.CreateTime)
		rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "ArticleTotal", total)
	}
}
