package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/util"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

// 查询所有标签，按创建时间排序
func QueryAllTagsOrderByTime(resp *[]Tag) error {
	return DB.Order("create_time").Find(&resp).Error
}

func HasTagByID(tagID int) (*Tag, bool) {
	tag := new(Tag)
	DB.Where("id = ?", tagID).First(&tag)
	if tag.ID == 0 {
		return nil, false
	}
	return tag, true
}

func AddTag(name string) error {
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		log.Printf("AddTag Begin Error, Name = [%v], : [%v]\n", name, err)
		return err
	}
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("AddTag Error, Callbacked... tagName = [%v]; [%v]", name, err)
		}
		tx.Commit()
	}()

	err = tx.Create(&Tag{
		Name:       name,
		CreateTime: time.Now(),
	}).Error
	return err
}

func insertTagToCache(tag *Tag) error {
	rc := RedisPool.Get()
	defer rc.Close()
	tagInfoCacheFields := []string{
		"ID", "Name", "CreateTime",
	}
	value := reflect.ValueOf(tag).Elem()
	for _, field := range tagInfoCacheFields {
		err := rc.Send("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID),
			field, value.FieldByName(field).String())
		if err != nil {
			msg := fmt.Sprintf("send fail when insert tag to cache, tagID = %v, field = %v", tag.ID, field)
			util.ExceptionLog(err, msg)
			return err
		}
	}
	err := rc.Flush()
	util.ExceptionLog(err, "flush fail when insert tag to cache")
	return err
}

func queryTagByIDFromDB(id int) (*Tag, error) {
	result := Tag{}
	err := DB.Where("id = ?", id).First(&result).Error
	return &result, err
}

func queryTagByIDFromCache(id int) (*Tag, error) {
	rc := RedisPool.Get()
	defer rc.Close()
	tagInfoCacheFields := []string{
		"ID", "Name", "CreateTime",
	}
	cacheReturnResults := make([]string, len(tagInfoCacheFields))
	for _, field := range tagInfoCacheFields {
		err := rc.Send("HGET", consts.TagsInfoHashCache+strconv.Itoa(id), field)
		if err != nil {
			msg := fmt.Sprintf("Fail to query tag info from cache, tagID = %v, field = %v", id, field)
			util.ExceptionLog(err, msg)
			return nil, err
		}
	}
	util.ExceptionLog(rc.Flush(), "Fail to flush query tag info")
	for i, field := range tagInfoCacheFields {
		result, err := rc.Receive()
		if err != nil {
			msg := fmt.Sprintf("Fail to do rc.Receive(), tagID = %v, field = %v", id, field)
			util.ExceptionLog(err, msg)
			return nil, err
		}
		if result != nil {
			cacheReturnResults[i] = string(result.([]byte))
		} else {
			// 缓存失效
			msg := fmt.Sprintf("缓存未命中！tagID = %v, field = %v", id, field)
			util.LogPlus(msg)
			tagFromDB, err := queryTagByIDFromDB(id)
			if err != nil {
				msg := fmt.Sprintf("Fail to query tag from DB when cache miss, tagID = %v", id)
				util.ExceptionLog(err, msg)
				return nil, err
			}
			err = insertTagToCache(tagFromDB)
			if err != nil {
				return nil, err
			}
			return tagFromDB, nil
		}
	}
	tagID, _ := strconv.Atoi(cacheReturnResults[0])
	createTime, _ := time.Parse("2006-01-02 15:04:05 +0000 UTC", cacheReturnResults[2])
	return &Tag{
		ID:         tagID,
		Name:       cacheReturnResults[1],
		CreateTime: createTime,
	}, nil
}

func QueryTagByID(id int) (*Tag, error) {
	if src.Setting.Redis {
		return queryTagByIDFromCache(id)
	}
	return QueryTagByID(id)
}
