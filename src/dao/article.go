package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/util"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var fields = []string{
	"Title", "Abstract", "ID", "AuthorID", "CreateTime", "Tags",
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
func setNewArticleInfoToCache(id int, article *ArticleListInfo) error {
	rc := RedisPool.Get()
	defer rc.Close()
	fields := []string{
		"Title", "Abstract", "ID", "Author", "CreateTime",
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
	// 插入Tag 信息
	tagIDs := make([]string, len(article.Tags))
	for i, tag := range article.Tags {
		tagIDs[i] = strconv.Itoa(tag.ID)
	}
	_, err := rc.Do("HSET", consts.ArticleInfoHashCache+strconv.Itoa(id),
		"Tags", strings.Join(tagIDs, consts.CacheTagsSplitStr))
	if err != nil {
		msg := fmt.Sprintf("Failed to insert tag information into cache, tags = %v, articleID = %v", tagIDs, id)
		util.ExceptionLog(err, msg)
		return err
	}
	return nil
}

/**
* WiKi: 文章信息未命中缓存时更新缓存
* Author: JuneBao
* Time: 2020/8/22 11:59
**/
func articleInfoMissHitsUpdate(id int) (ArticleListInfo, error) {
	ali, err := QueryArticleListInfoByIDWithDB(id)
	if err != nil {
		log.Printf("缓存未命中后从数据库中获取信息失败！id = [%v]", id)
		return ali, err
	}
	if ali.ID == 0 {
		log.Printf("缓存未命中，id对应文章信息不存在！id = [%v]", id)
		return ali, err
	}
	err = setNewArticleInfoToCache(id, &ali)
	if err != nil {
		log.Printf("缓存未命中后更新缓存失败， id = [%v]", id)
		return ali, err
	}
	return ali, nil
}

/**
* WiKi: 从缓存中查询单个文章的信息,没有查询到会从数据库中补充
* Author: JuneBao
* Time: 2020/8/22 16:14
**/
func queryArticleInfoFromCache(id int) (ArticleListInfo, error) {
	result := ArticleListInfo{}
	articleFields := make([]string, 0)
	rc := RedisPool.Get()
	tags := make([]Tag, 0)
	// 如果这篇文章的标签信息没在缓存中， 就会从数据库中查找，这样就不需要再根据id查找了
	queryTagFlag := true
	defer rc.Close()
	for _, field := range fields {
		err := rc.Send("Hget", consts.ArticleInfoHashCache+strconv.Itoa(id), field)
		if err != nil {
			msg := fmt.Sprintf("Hget article info from redis fail, article id = %v, field = %v", id, field)
			util.ExceptionLog(err, msg)
		}
	}
	util.ExceptionLog(rc.Flush(), "redis flush fail")
	for _, field := range fields {
		r, err := rc.Receive()
		if err != nil {
			msg := fmt.Sprintf("do rc.Receive fail, article id = %v", id)
			util.ExceptionLog(err, msg)
			return result, err
		}
		// 缓存未命中
		if r == nil {
			log.Printf("缓存未命中！article id = %v, field = %v\n", id, field)
			a, _ := articleInfoMissHitsUpdate(id)
			value := reflect.ValueOf(a).FieldByName(field).String()
			if field == "Tags" {
				tagStr := make([]string, len(a.Tags))
				for i, tag := range a.Tags {
					tagStr[i] = strconv.Itoa(tag.ID)
				}
				articleFields = append(articleFields, strings.Join(tagStr, consts.CacheTagsSplitStr))
				tags = a.Tags
				queryTagFlag = false
			} else {
				articleFields = append(articleFields, value)
			}
			continue
		}
		articleFields = append(articleFields, string(r.([]byte)))
	}
	if len(articleFields) >= len(fields) {
		id, _ := strconv.Atoi(articleFields[2])
		cTime, _ := strconv.Atoi(articleFields[4])
		tagIDs := strings.Split(articleFields[5], "-")
		if queryTagFlag {
			for _, tagID := range tagIDs {
				tid, _ := strconv.Atoi(tagID)
				tag, err := QueryTagByID(tid)
				if err != nil {
					return result, nil
				}
				tags = append(tags, *tag)
			}
		}
		result.ID = id
		result.Title = articleFields[0]
		result.Author = "Junebao"
		result.Abstract = articleFields[1]
		result.CreateTime = int64(cTime)
		result.Tags = tags
	}
	return result, nil
}

/**
* WiKi: 通过缓存查询文章列表页信息
* Author: JuneBao
* Time: 2020/9/11 23:27
**/
func queryArticleInfoByLimitByCache(page, pageSize, total int) ([]ArticleListInfo, error) {
	result := make([]ArticleListInfo, 0)
	// 获取 ArticleID List
	// startTime := time.Now().UnixNano()
	ids, err := queryArticleIDListFromCache(page, pageSize, total)
	// step1Time := time.Now().UnixNano()
	// log.Printf("查询文章ID List 用时：%v\n", (step1Time - startTime) / 1000000)
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
		var a ArticleListInfo
		var e error
		// step2 := time.Now().UnixNano()
		a, e = queryArticleInfoFromCache(id)
		// step3 := time.Now().UnixNano()
		// log.Printf("queryArticleInfoFromCache 用时：%v\n", (step3 - step2) / 1000000)
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
func QueryArticleInfoByLimit(page, pageSize int) ([]ArticleListInfo, int, error) {
	// startTime := time.Now().UnixNano()
	total, err := QueryArticleTotal()
	if err != nil {
		return nil, 0, err
	}
	// step1Time := time.Now().UnixNano()
	// log.Printf("查询文章总数用时：%v\n", (step1Time - startTime) / 1000000)
	if src.Setting.Redis {
		result, err := queryArticleInfoByLimitByCache(page, pageSize, total)
		// step2Time := time.Now().UnixNano()
		// log.Printf("查询文章列表用时：%v\n", (step2Time - step1Time) / 1000000)
		if err == nil {
			return result, total, err
		}
	}

	start := (page - 1) * pageSize
	if total < start {
		util.LogPlus("total < start")
		return nil, 0, errors.New("total < start")
	}
	newPageSize := total - start
	if newPageSize < pageSize {
		pageSize = newPageSize
	}
	articleList := make([]Article, 0)

	err = DB.Order("create_time desc").Limit(pageSize).Offset(start).Find(&articleList).Error
	if err != nil {
		return nil, 0, err
	}
	articleListInfos := make([]ArticleListInfo, len(articleList))
	for i, article := range articleList {
		tags := make([]Tag, 0)
		err := QueryAllTagsByArticleID(article.ID, &tags)
		if err != nil {
			msg := fmt.Sprintf("get article tags fail, article id = %v", article.ID)
			util.ExceptionLog(err, msg)
		}
		articleListInfos[i] = ArticleListInfo{
			Tags:       tags,
			ID:         article.ID,
			Title:      article.Title,
			CreateTime: article.CreateTime.Unix(),
			Abstract:   article.Abstract,
			Author:     "Junebao",
		}
	}
	return articleListInfos, total, nil
}

/**
* WiKi: 查询某个 Tag 下的单页文章信息
* Author: JuneBao
* Time: 2020/9/24 10:56
**/
func QueryArticleInfoByLimitWithTag(tagID, page, pageSize int) ([]ArticleListInfo, int, error) {
	aIDs, err := QueryAllArticleByTagID(tagID)
	if err != nil {
		return nil, 0, err
	}

	total := len(aIDs)
	start := (page - 1) * pageSize
	log.Printf("total = %v, pageSize = %v", total, pageSize)
	if total < start {
		util.LogPlus("total < start")
		return nil, 0, errors.New("total < start")
	}
	newPageSize := total - start
	if newPageSize < pageSize {
		pageSize = newPageSize
	}

	articleList := make([]ArticleListInfo, 0)

	if total > pageSize {
		aIDs = aIDs[start : start+pageSize]
	}

	for _, a := range aIDs {
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
func QueryArticleListInfoByIDWithDB(id int) (ArticleListInfo, error) {
	result := ArticleListInfo{}
	article := Article{}
	err := DB.Select("id, title, abstract,"+
		" author_id, create_time").Where("id = ?", id).First(&article).Error
	if err != nil {
		msg := fmt.Sprintf("query article by id fail, article id = %v", id)
		util.ExceptionLog(err, msg)
		return result, err
	}
	tags := make([]Tag, 0)
	err = QueryAllTagsByArticleID(id, &tags)
	if err != nil {
		msg := fmt.Sprintf("query all tags by articleID fail, articleID = %v", id)
		util.ExceptionLog(err, msg)
		return result, err
	}
	result = ArticleListInfo{
		ID:         article.ID,
		Title:      article.Title,
		Abstract:   article.Abstract,
		CreateTime: article.CreateTime.Unix(),
		Author:     "Junebao",
		Tags:       tags,
	}
	return result, nil
}

/**
* WiKi: 通过 ID 查询文章信息
* Author: JuneBao
* Time: 2020/9/24 10:56
**/
func QueryArticleByID(id int) (ArticleListInfo, error) {
	var articleListInfo ArticleListInfo
	if src.Setting.Redis {
		return queryArticleInfoFromCache(id)
	}
	articleListInfo, err := QueryArticleListInfoByIDWithDB(id)

	return articleListInfo, err
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

func addArticleWithCache(newArticle *Article, tags []Tag) error {
	// 更新 ArticleIDList
	rp := RedisPool.Get()
	var err error
	defer func() {
		rp.Close()
	}()
	re, err := rp.Do("LPUSH", consts.ArticleIDListCache, newArticle.ID)
	if err != nil {
		msg := fmt.Sprintf("update %v fail, article id = %v", consts.ArticleIDListCache, newArticle.ID)
		util.ExceptionLog(err, msg)
		return err
	}

	// 更新 ArticleListInfo
	err = setNewArticleInfoToCache(int(re.(int64)), &ArticleListInfo{
		ID:         newArticle.ID,
		Title:      newArticle.Title,
		Abstract:   newArticle.Abstract,
		CreateTime: newArticle.CreateTime.Unix(),
		Author:     "Junebao",
		Tags:       tags,
	})
	if err != nil {
		msg := fmt.Sprintf("update %v fail, article id = %v", "article info", newArticle.ID)
		util.ExceptionLog(err, msg)
	}
	return err
}

func AddArticle(newArticle *Article, tagIDs []int) (*Article, error) {
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
	// 更新 article_tag 表
	tags := make([]Tag, len(tagIDs))
	for i, tagID := range tagIDs {
		tag, err := QueryTagByID(tagID)
		if err != nil {
			return nil, err
		}
		if tag == nil {
			msg := fmt.Sprintf("Get nil when query tag, tagID = %v", tagID)
			err := errors.New("return nil")
			util.ExceptionLog(err, msg)
			return nil, err
		}
		tags[i] = *tag
		err = InsertArticleTag(&ArticleTags{
			ArticleID: newArticle.ID,
			TagID:     tagID,
		})
		if err != nil {
			msg := fmt.Sprintf("fail to insert new article tag;"+
				" articleID = %v, tagID = %v", newArticle.ID, tagID)
			util.ExceptionLog(err, msg)
			return nil, err
		}
	}
	if src.Setting.Redis {
		_ = addArticleWithCache(newArticle, tags)
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
	if err != nil {
		msg := fmt.Sprintf("Fail to update article table, articleID = %v", id)
		util.ExceptionLog(err, msg)
		return err
	}
	if src.Setting.Redis {
		err = updateArticleWithCache(id, article)
		if err != nil {
			msg := fmt.Sprintf("Fail to update article cache, articleID = %v", id)
			util.ExceptionLog(err, msg)
			return err
		}
	}
	return nil
}
