package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/util"
	"fmt"
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
func queryArticleIDListFromCache(page, pageSize, total int) ([]int, error) {
	start := (page-1)*pageSize + 1
	r := make([]int, 0)
	rc := RedisPool.Get()
	defer rc.Close()
	newPageSize := total - start + 1
	if newPageSize < pageSize {
		pageSize = newPageSize
	}
	for i := 0; i < pageSize; i++ {
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
	immutable := reflect.ValueOf(article).Elem()
	for _, field := range fields {
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
	err := QueryArticleByIDWithDB(id, &articleDB)
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
* WiKi: 从缓存中查询单个文章的信息,没有查询到会从数据库中补充
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
		if err != nil {
			mes := fmt.Sprintf("query article info from cache fail, article id = %v", id)
			util.ExceptionLog(err, mes)
		}
	}
	util.ExceptionLog(rc.Flush(), "redis flush fail")
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

/**
* WiKi: 通过缓存查询文章列表页信息
* Author: JuneBao
* Time: 2020/9/11 23:27
**/
func queryArticleInfoByLimitByCache(page, pageSize, total int) ([]Article, error) {
	result := make([]Article, 0)
	// 获取 ArticleID List
	ids, err := queryArticleIDListFromCache(page, pageSize, total)
	if err != nil {
		log.Printf("从缓存中获取文章ID列表失败")
		// 获取列表失败，开启携程重置数据
		go func() {
			iErr := InitArticleIDListCache()
			if iErr != nil {
				msg := fmt.Sprintf("Failed to update the article ID in the cache asynchronously")
				util.ExceptionLog(iErr, msg)
			}
		}()
		return nil, err
	}
	for _, id := range ids {
		var a Article
		var e error
		a, e = queryArticleInfoFromCache(id)
		if e != nil {
			return nil, e
		}
		result = append(result, a)
	}

	return result, nil
}

/**
* WiKi: 查询单页文章列表信息
* Author: JuneBao
* Time: 2020/9/11 23:27
**/
func QueryArticleInfoByLimit(page, pageSize, total int) ([]Article, error) {

	if src.Setting.Redis {
		result, err := queryArticleInfoByLimitByCache(page, pageSize, total)
		if err == nil {
			return result, err
		}
	}
	start := (page - 1) * pageSize
	newPageSize := total - start
	if newPageSize < pageSize {
		pageSize = newPageSize
	}
	articleList := make([]Article, 0)
	al := DB.Order("create_time desc").Limit(pageSize).Offset(start).Find(&articleList)
	return articleList, al.Error
}

/**
* WiKi: 查询某个 Tag 下的单页文章信息
* Author: JuneBao
* Time: 2020/9/24 10:56
**/
func QueryArticleInfoByLimitWithTag(tagID, page, pageSize int) ([]Article, int, error) {
	articleTags, err := QueryAllArticleByTagID(tagID)
	if err != nil {
		return nil, 0, err
	}

	total := len(articleTags)
	start := (page - 1) * pageSize
	newPageSize := total - start
	if newPageSize < pageSize {
		pageSize = newPageSize
	}

	articleList := make([]Article, 0)
	log.Printf("total = %v, pageSize = %v", total, pageSize)
	if total > pageSize {
		articleTags = articleTags[start : start+pageSize]
	}

	for _, a := range articleTags {
		article, err := QueryArticleByID(a.ArticleID)
		if err != nil {
			return nil, 0, err
		}
		articleList = append(articleList, article)
	}
	return articleList, total, nil
}

func QueryAllArticle(articleList *[]Article) error {
	return DB.Order("create_time desc").Find(&articleList).Error
}

/**
* WiKi: 通过文章ID在数据库中查询文章信息
* Author: JuneBao
* Time: 2020/8/22 11:40
**/
func QueryArticleByIDWithDB(id int, article *Article) error {
	return DB.Select("id, title, abstract,"+
		" author_id, create_time").Where("id = ?", id).First(&article).Error
}

/**
* WiKi: 通过 ID 查询文章信息
* Author: JuneBao
* Time: 2020/9/24 10:56
**/
func QueryArticleByID(id int) (Article, error) {
	var article Article
	if src.Setting.Redis {
		return queryArticleInfoFromCache(id)
	}
	err := QueryArticleByIDWithDB(id, &article)
	return article, err
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

func addArticleWithCache(newArticle *Article) error {
	// 更新 ArticleIDList
	rp := RedisPool.Get()
	var err error
	defer func() {
		rp.Close()
	}()
	re, err := rp.Do("RPUSH", consts.ArticleIDListCache, newArticle.ID)
	log.Println(re)
	if err != nil {
		msg := fmt.Sprintf("update %v fail, article id = %v", consts.ArticleIDListCache, newArticle.ID)
		util.ExceptionLog(err, msg)
		return err
	}

	// 更新 ArticleListInfo
	err = setNewArticleInfoToCache(int(re.(int64)), newArticle)
	if err != nil {
		msg := fmt.Sprintf("update %v fail, article id = %v", "article info", newArticle.ID)
		util.ExceptionLog(err, msg)
	}
	return err
}

func AddArticle(newArticle *Article) (*Article, error) {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("insert new article fail, title = %v", newArticle.Title)
			util.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Create(newArticle).Error
	if src.Setting.Redis {
		_ = addArticleWithCache(newArticle)
	}
	return newArticle, err
}

// TODO: bitmap
func HasArticle(id int) bool {
	return true
}

func updateArticleWithCache(id int, article *Article) error {
	rc := RedisPool.Get()
	defer func() {
		rc.Close()
	}()
	value := reflect.ValueOf(article).Elem()
	for _, field := range fields {
		err := rc.Send("HSet", consts.ArticleInfoHashCache+strconv.Itoa(id),
			field, value.FieldByName(field))
		if err != nil {
			msg := fmt.Sprintf("send fail, id = %v, field = %v", id, field)
			util.ExceptionLog(err, msg)
			return err
		}
	}
	err := rc.Flush()
	if err != nil {
		msg := fmt.Sprintf("flush fail, id = %v", id)
		util.ExceptionLog(err, msg)
		return err
	}
	return nil
}

func UpdateArticle(id int, article *Article) error {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("update article fail, id = %v, title = %v", id, article.Title)
			util.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Model(&Article{}).Where("id = ?", id).Updates(article).Error
	err = updateArticleWithCache(id, article)
	return err
}
