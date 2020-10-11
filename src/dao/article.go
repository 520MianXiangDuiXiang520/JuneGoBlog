package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/junebao.top/utils"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
	"strconv"
	"strings"
)

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

		if result == nil {
			return nil, errors.New("ReceiveNil")
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
		utils.ExceptionLog(err, msg)
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
	var fields = []string{
		"Title", "Abstract", "ID", "CreateTime", "Tags",
	}
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
			utils.ExceptionLog(err, msg)
		}
	}
	utils.ExceptionLog(rc.Flush(), "redis flush fail")
	for _, field := range fields {
		r, err := rc.Receive()
		if err != nil {
			msg := fmt.Sprintf("do rc.Receive fail, article id = %v", id)
			utils.ExceptionLog(err, msg)
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
		cTime, _ := strconv.Atoi(articleFields[3])
		tagIDs := strings.Split(articleFields[4], "-")
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

	ids, err := queryArticleIDListFromCache(page, pageSize, total)

	if err != nil {
		utils.LogPlus("从缓存中获取文章ID列表失败")
		go func() {
			iErr := InitArticleIDListCache()
			if iErr != nil {
				msg := fmt.Sprintf("Failed to update the article ID in the cache asynchronously")
				utils.ExceptionLog(iErr, msg)
			}
		}()
		return nil, err
	}
	for _, id := range ids {
		var a ArticleListInfo
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
func QueryArticleInfoByLimit(page, pageSize int) ([]ArticleListInfo, int, error) {
	total, err := QueryArticleTotal()
	if err != nil {
		return nil, 0, err
	}
	if src.Setting.Redis {
		result, err := queryArticleInfoByLimitByCache(page, pageSize, total)
		if err == nil {
			return result, total, err
		}
		utils.LogPlus("Fail to query from cache!!")
	}

	// 缓存中没有查到，从数据库中查
	start := (page - 1) * pageSize
	if total < start {
		utils.LogPlus("total < start")
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
			utils.ExceptionLog(err, msg)
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
	if total < start {
		utils.LogPlus("total < start")
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
		utils.ExceptionLog(err, msg)
		return result, err
	}
	tags := make([]Tag, 0)
	err = QueryAllTagsByArticleID(id, &tags)
	if err != nil {
		msg := fmt.Sprintf("query all tags by articleID fail, articleID = %v", id)
		utils.ExceptionLog(err, msg)
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
	result, err := redis.Int(rc.Do("LLen", consts.ArticleIDListCache))
	if err != nil {
		msg := fmt.Sprintf("Fail to query articles total")
		utils.ExceptionLog(err, msg)
		return 0, err
	}
	if result == 0 {
		msg := fmt.Sprintf("The total number of articles queried from the cache is 0")
		err := errors.New("ResultIsZero")
		utils.ExceptionLog(err, msg)
		return 0, err
	}
	return result, nil
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
		utils.ExceptionLog(err, msg)
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
		utils.ExceptionLog(err, msg)
	}
	return err
}

func AddArticle(newArticle *Article, tagIDs []int) (*Article, error) {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("insert new article fail, title = %v", newArticle.Title)
			utils.ExceptionLog(err, msg)
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
			utils.ExceptionLog(err, msg)
			return nil, err
		}
		tags[i] = *tag
		// 不管什么地方发生错误，立刻回滚
		err = tx.Create(&ArticleTags{
			ArticleID: newArticle.ID,
			TagID:     tagID,
		}).Error
		if err != nil {
			msg := fmt.Sprintf("fail to insert new article tag;"+
				" articleID = %v, tagID = %v", newArticle.ID, tagID)
			utils.ExceptionLog(err, msg)
			return nil, err
		}
	}
	if src.Setting.Redis {
		_ = addArticleWithCache(newArticle, tags)
	}
	return newArticle, err
}

func hasArticleWithDB(id int) bool {
	article := &Article{}
	DB.Where("id = ?", id).First(article)
	return article.ID != 0
}

func HasArticle(id int) bool {
	if src.Setting.Redis {
		// TODO: bitmap
	}
	return hasArticleWithDB(id)
}

func updateArticleWithCache(id int, article *Article) error {
	rc := RedisPool.Get()
	defer func() {
		rc.Close()
	}()
	var err error
	err = rc.Send("HSet", consts.ArticleInfoHashCache+strconv.Itoa(id),
		"Title", article.Title)
	err = rc.Send("HSet", consts.ArticleInfoHashCache+strconv.Itoa(id),
		"Abstract", article.Abstract)
	if err != nil {
		msg := fmt.Sprintf("Fail to send new value with cache when update article, article = %v", article)
		utils.ExceptionLog(err, msg)
		return err
	}
	err = rc.Flush()
	if err != nil {
		msg := fmt.Sprintf("flush fail, id = %v", id)
		utils.ExceptionLog(err, msg)
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
			utils.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	// 更新文章不更新创建时间
	err = tx.Model(&Article{}).Omit("create_time").Where("id = ?", id).Updates(article).Error
	if err != nil {
		msg := fmt.Sprintf("Fail to update article table, articleID = %v", id)
		utils.ExceptionLog(err, msg)
		return err
	}
	if src.Setting.Redis {
		err = updateArticleWithCache(id, article)
		if err != nil {
			msg := fmt.Sprintf("Fail to update article cache, articleID = %v", id)
			utils.ExceptionLog(err, msg)
			return err
		}
	}
	return nil
}

func deleteArticleFromDB(id int) error {
	DB.LogMode(true)
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("Fail to delete article from db, article id = %v", id)
			utils.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	// 删除 article 表中的数据
	err = tx.Where("id = ?", id).Delete(&Article{}).Error
	// 删除 article_tag 表中的数据
	err = tx.Where("article_id = ?", id).Delete(&ArticleTags{}).Error
	return err
}

func deleteArticleIDListCacheByID(id int) error {
	rc := RedisPool.Get()
	defer rc.Close()

	_, err := rc.Do("LREM", consts.ArticleIDListCache, 0, id)
	if err != nil {
		msg := fmt.Sprintf("Fail to delete ArticleIDListCache by id, id = %v", id)
		utils.ExceptionLog(err, msg)
		return err
	}
	return nil
}

func deleteArticleInfoHashCacheByID(id int) error {
	rc := RedisPool.Get()
	defer rc.Close()

	_, err := rc.Do("DEL", consts.ArticleInfoHashCache+strconv.Itoa(id))
	if err != nil {
		msg := fmt.Sprintf("Fail to delete ArticleInfoHashCache by id, id = %v", id)
		utils.ExceptionLog(err, msg)
		return err
	}
	return nil
}

func deleteArticleFromCache(id int) error {
	// 修改 ArticleIDListCache
	err := deleteArticleIDListCacheByID(id)
	if err != nil {
		return err
	}
	// 删除 ArticleInfoHashCache
	err = deleteArticleInfoHashCacheByID(id)
	if err != nil {
		return err
	}

	// TODO: 修改 BitMap
	return nil
}

func DeleteArticle(id int) error {
	// 1. 从数据库中删除
	err := deleteArticleFromDB(id)
	if err != nil {
		return err
	}
	// 2. 从 Redis 中删除
	if src.Setting.Redis {
		err := deleteArticleFromCache(id)
		if err != nil {
			return err
		}
	}
	return nil
}
