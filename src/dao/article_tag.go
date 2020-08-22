package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/util"
	"log"
	"strconv"
)

func QueryArticleTotalByTagIDFromCache(tagID int) (int, error) {
	rc := RedisPool.Get()
	defer rc.Close()
	r, e := rc.Do("Hget", consts.TagsInfoHashCache+strconv.Itoa(tagID), "ArticleTotal")
	if e != nil {
		log.Printf("QueryArticleTotalByTagIDFromCache 执行失败， tagId = [%v]", tagID)
		return 0, e
	}
	return strconv.Atoi(string(r.([]byte)))
}

func QueryAllTagsByArticleID(articleID int, tags *[]Tag) error {
	at := make([]ArticleTags, 0)
	DB.Where("article_id = ?", articleID).Find(&at)
	tagsID := make([]int, 0)
	for _, tag := range at {
		tagsID = append(tagsID, tag.TagID)
	}
	return DB.Where("id IN (?)", tagsID).Find(&tags).Error
}

func QueryArticleTotalByTagIDFromDB(tagID int) int {
	var total int
	DB.Model(&ArticleTags{}).Where("tag_id = ?", tagID).Count(&total)
	log.Println(total)
	return total
}

func QueryArticleTotalByTagID(tagID int) int {
	if src.Setting.Redis {
		r, e := QueryArticleTotalByTagIDFromCache(tagID)
		util.CatchException(e)
		return r
	}
	return QueryArticleTotalByTagIDFromDB(tagID)
}
