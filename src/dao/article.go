package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/util"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
	"strconv"
	"time"
)

var fields = []string{
	"Title", "Abstract", "ID", "AuthorID", "CreateTime",
}

/**
* WiKi: 通过缓存中文章ID列表的下标得到文章ID
* Author: JuneBao
* Time: 2020/8/22 16:21
**/
func QueryArticleIDFromCacheByIndex(index int) (int, error) {
	rc := RedisPool.Get()
	defer rc.Close()
	value, err := redis.String(rc.Do("LINDEX", consts.ArticleIDListCache, index))
	v, _ := strconv.Atoi(value)
	return v, err
}

/**
* WiKi: 从缓存中查询文章ID列表
* Author: JuneBao
* Time: 2020/8/22 0:45
**/
func queryArticleIDListFromCache(page, pageSize int) ([]int, error) {
	start := (page-1)*pageSize + 1
	r := make([]int, 0)
	rc := RedisPool.Get()
	defer rc.Close()
	total, err := QueryArticleTotal()
	if err != nil {
		return nil, err
	}
	pageSize = total - start + 1
	for i := 0; i < pageSize; i++ {
		//log.Printf("Do Redis: LINDEX %v %v", consts.ArticleIDListCache, i + start - 1)
		if err := rc.Send("lIndex", consts.ArticleIDListCache, i+start-1); err != nil {
			log.Fatal("Send Lindex ERROR!")
			return r, err
		}
	}
	if err := rc.Flush(); err != nil {
		log.Fatal("Flush Lindex ERROR!")
		return r, err
	}

	for i := 0; i < pageSize; i++ {
		result, err := rc.Receive()
		if err != nil {
			log.Fatal("Send Lindex ERROR!")
			return r, err
		}

		re, _ := strconv.Atoi(string(result.([]byte)))
		r = append(r, re)
	}
	return r, nil

}

/**
* WiKi: 向缓存中插入一条文章信息记录
* Author: JuneBao
* Time: 2020/8/22 11:52
**/
func setNewArticleInfoToCache(id int, article *Article) error {
	rc := RedisPool.Get()
	defer rc.Close()
	fields := []string{
		"Title", "Abstract", "ID", "AuthorID", "CreateTime",
	}
	for _, field := range fields {
		immutable := reflect.ValueOf(article)
		val := immutable.FieldByName(field)
		_, err := rc.Do("HSET", consts.ArticleInfoHashCache+strconv.Itoa(id), field, val)
		if err != nil {
			log.Printf("setNewArticleInfoToCache 执行 HSET 失败， id = [%v]", id)
			return err
		}
	}
	return nil
}

/**
* WiKi: 文章信息未命中缓存时更新缓存
* Author: JuneBao
* Time: 2020/8/22 11:59
**/
func articleInfoMissHitsUpdate(id int) (Article, error) {
	articleDB := Article{}
	err := QueryArticleInfoByID(id, &articleDB)
	if err != nil {
		log.Printf("缓存未命中后从数据库中获取信息失败！id = [%v]", id)
		return articleDB, err
	}
	if articleDB.ID == 0 {
		log.Printf("缓存未命中，id对应文章信息不存在！id = [%v]", id)
		return articleDB, err
	}
	err = setNewArticleInfoToCache(id, &articleDB)
	if err != nil {
		log.Printf("缓存未命中后更新缓存失败， id = [%v]", id)
		return articleDB, err
	}
	return articleDB, nil
}

/**
* WiKi: 从缓存中获取文章信息列表
* Input: 要查询的文章列表
* Author: JuneBao
* Time: 2020/8/22 2:05
**/
func getArticleInfoListFromCache(ids []int) ([]Article, error) {
	result := make([]Article, 0)
	var err error
	for _, id := range ids {
		a, e := queryArticleInfoFromCache(id)
		if e != nil {
			log.Printf("getArticleInfoListFromCache Call queryArticleInfoFromCache Error")
			return result, err
		}
		result = append(result, a)

	}
	return result, nil
}

/**
* WiKi: 从缓存中查询单个文章的信息
* Author: JuneBao
* Time: 2020/8/22 16:14
**/
func queryArticleInfoFromCache(id int) (Article, error) {
	result := Article{}
	articleFields := make([]string, 0)
	rc := RedisPool.Get()
	defer rc.Close()
	for _, field := range fields {
		err := rc.Send("Hget", consts.ArticleInfoHashCache+strconv.Itoa(id), field)
		util.CatchException(err)
	}
	util.CatchException(rc.Flush())
	for range fields {
		r, err := rc.Receive()
		if err != nil {
			log.Println("queryArticleInfoFromCache Send HGet Error")
			return result, err
		}
		// 缓存未命中
		if r == nil {
			a, _ := articleInfoMissHitsUpdate(id)
			articleFields = append(articleFields, a.Title)
			continue
		}
		articleFields = append(articleFields, string(r.([]byte)))
	}

	if len(articleFields) >= len(fields) {
		authorID, _ := strconv.Atoi(articleFields[3])
		id, _ := strconv.Atoi(articleFields[2])
		cTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", articleFields[4])
		result.ID = id
		result.Title = articleFields[0]
		result.AuthorID = authorID
		result.Abstract = articleFields[1]
		result.CreateTime = cTime
		result.Text = ""
	}
	return result, nil

}

func QueryArticleByLimit(page, pageSize int) ([]Article, error) {
	if src.Setting.Redis {
		ids, err := queryArticleIDListFromCache(page, pageSize)
		if err != nil {
			log.Printf("从缓存中获取文章ID列表失败")
			return nil, err
		}
		res, err := getArticleInfoListFromCache(ids)
		return res, err
	}
	articleList := make([]Article, 0)
	al := DB.Limit(pageSize).Offset(page * pageSize).Find(&articleList)
	return articleList, al.Error
}

func QueryAllArticle(articleList *[]Article) error {
	return DB.Find(&articleList).Error
}

/**
* WiKi: 通过文章ID在数据库中查询文章信息
* Author: JuneBao
* Time: 2020/8/22 11:40
**/
func QueryArticleInfoByID(id int, article *Article) error {
	return DB.Where("id = ?", id).First(&article).Error
}

func HasArticle(id int) bool {
	a := Article{}
	DB.Where("id = ?", id).First(&a)
	return a.ID != 0
}

func QueryArticleDetail(id int) (Article, error) {
	a := Article{}
	r := DB.Where("id = ?", id).First(&a)
	return a, r.Error
}

func queryArticleTotalByCache() (int, error) {
	rc := RedisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do("LLen", consts.ArticleIDListCache))
}

func queryArticleTotalByDB() (int, error) {
	var total int
	c := DB.Model(&Article{}).Count(&total)
	return total, c.Error
}

func QueryArticleTotal() (int, error) {
	var total int
	var err error
	if src.Setting.Redis {
		total, err = queryArticleTotalByCache()
		if err != nil {
			log.Printf("通过缓存获取文章总数失败！！！: %v", err)
			return queryArticleTotalByDB()
		}
		return total, nil
	}
	return queryArticleTotalByDB()
}
