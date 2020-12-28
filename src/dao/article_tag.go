package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"fmt"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"log"
	"strconv"
	"strings"
)

// 暂不使用缓存
// func QueryArticleTotalByTagIDFromCache(tagID int) (int, error) {
// 	rc := RedisPool.Get()
// 	defer rc.Close()
// 	r, e := rc.Do("Hget", consts.TagsInfoHashCache+strconv.Itoa(tagID), "ArticleTotal")
// 	if e != nil {
// 		log.Printf("QueryArticleTotalByTagIDFromCache 执行失败， tagId = [%v]", tagID)
// 		return 0, e
// 	}
// 	return strconv.Atoi(string(r.([]byte)))
// }

func InsertArticleTag(at *ArticleTags) error {
	tx := juneDao.GetDB().Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("insert  articleTag fail, article id = %v, tag id = %v,", at.ArticleID, at.TagID)
			juneLog.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Create(at).Error
	return err
}

func QueryAllTagsByArticleID(articleID int, tags *[]Tag) error {
	at := make([]ArticleTags, 0)
	juneDao.GetDB().Where("article_id = ?", articleID).Find(&at)
	tagsID := make([]int, 0)
	for _, tag := range at {
		tagsID = append(tagsID, tag.TagID)
	}
	return juneDao.GetDB().Where("id IN (?)", tagsID).Find(&tags).Error
}

func QueryArticleTotalByTagIDFromDB(tagID int) int {
	var total int
	juneDao.GetDB().Model(&ArticleTags{}).Where("tag_id = ?", tagID).Count(&total)
	return total
}

func QueryArticleTotalByTagID(tagID int) (int, error) {
	// var err error
	// var result int
	// if src.GetSetting().Others.Redis {
	// 	result, err = QueryArticleTotalByTagIDFromCache(tagID)
	// 	if err != nil {
	// 		return QueryArticleTotalByTagIDFromDB(tagID), err
	// 	}
	// 	return result, nil
	// }
	return QueryArticleTotalByTagIDFromDB(tagID), nil
}

// 判断文章更新时 tags 是否发生了改变
func hasTagsChanged(articleID int, tags []*Tag) bool {
	history := make([]Tag, 0)
	err := QueryAllTagsByArticleID(articleID, &history)
	if err != nil {
		msg := fmt.Sprintf("query all tags by article id fail, article id = %v", articleID)
		juneLog.ExceptionLog(err, msg)
		return true
	}
	if len(history) != len(tags) {
		return true
	}
	for _, tag := range tags {
		get := false
		for _, his := range history {
			if tag.ID == his.ID {
				get = true
				break
			}
		}
		if !get {
			return true
		}
	}
	return false
}

func DeleteArticleTags(articleID int) error {
	tx := juneDao.GetDB().Begin()
	var err error
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("delete articleTag fail, article id = %v", articleID)
			juneLog.ExceptionLog(err, msg)
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Where("article_id = ?", articleID).Delete(&ArticleTags{}).Error
	return err
}

func updateArticleTagsToCache(articleID int, tags []*Tag) error {
	rc := juneDao.GetRedisConn()
	defer rc.Close()
	tIDs := make([]string, len(tags))
	for i, t := range tags {
		tIDs[i] = strconv.Itoa(t.ID)
	}
	_, err := rc.Do("HSET", consts.ArticleInfoHashCache+strconv.Itoa(articleID),
		"Tags", strings.Join(tIDs, consts.CacheTagsSplitStr))
	if err != nil {
		msg := fmt.Sprintf("do hset fail when update article tags, tIDs = %v", tIDs)
		juneLog.ExceptionLog(err, msg)
	}
	return err
}

func UpdateArticleTags(articleID int, tags []*Tag) error {
	// 如果 Tag 没有发生改变就不做修改
	if !hasTagsChanged(articleID, tags) {
		log.Printf("tag no change \n")
		return nil
	}
	var err error
	e := DeleteArticleTags(articleID)
	if e != nil {
		return e
	}
	for _, tag := range tags {
		err = InsertArticleTag(&ArticleTags{
			ArticleID: articleID,
			TagID:     tag.ID,
		})
		if err != nil {
			return err
		}
	}
	if src.GetSetting().Others.Redis {
		_ = updateArticleTagsToCache(articleID, tags)
	}
	return nil
}

func UpdateArticleTagsByIntList(articleID int, intTags []int) error {
	tags := make([]*Tag, len(intTags))
	for i, tID := range intTags {
		tags[i], _ = QueryTagByID(tID)
	}
	return UpdateArticleTags(articleID, tags)
}

/**
* WiKi: 查询 tagID 下的所有文章， 返回的是 ArticleTags 切片
* Author: JuneBao
* Time: 2020/9/24 10:58
**/
func QueryAllArticleByTagID(tagID int) ([]ArticleTags, error) {
	// TODO: 使用缓存
	result := make([]ArticleTags, 0)
	juneDao.GetDB().LogMode(true)
	err := juneDao.GetDB().Model(&ArticleTags{}).Where("tag_id = ?", tagID).Find(&result).Error

	if err != nil {
		msg := fmt.Sprintf("Query All Article By TagID fail, tagid = %v", tagID)
		juneLog.ExceptionLog(err, msg)
		return nil, err
	}
	return result, err
}
