package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/junebao.top/utils"
	"fmt"
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

func addTagFromDB(name string) (*Tag, error) {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			msg := fmt.Sprintf("Fail to add tag with DB, tag name = %v", name)
			utils.ExceptionLog(err, msg)
		}
		tx.Commit()
	}()
	newTag := Tag{
		Name:       name,
		CreateTime: time.Now(),
	}
	err = tx.Create(&newTag).Error
	return &newTag, err
}

func addTagFromCache(tag *Tag) error {
	rc := RedisPool.Get()
	var err error
	_, err = rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "ID", tag.ID)
	if err != nil {
		msg := fmt.Sprintf("Fail to Send TagsInfoHashCache:%v field = %v", tag.ID, "ID")
		utils.ExceptionLog(err, msg)
		return err
	}

	_, err = rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID), "Name", tag.Name)
	if err != nil {
		msg := fmt.Sprintf("Fail to Send TagsInfoHashCache:%v field = %v", tag.ID, "Name")
		utils.ExceptionLog(err, msg)
		return err
	}

	_, err = rc.Do("HSET", consts.TagsInfoHashCache+strconv.Itoa(tag.ID),
		"CreateTime", tag.CreateTime.Unix())
	if err != nil {
		msg := fmt.Sprintf("Fail to Send TagsInfoHashCache:%v field = %v", tag.ID, "CreateTime")
		utils.ExceptionLog(err, msg)
		return err
	}

	return nil
}

func AddTag(name string) error {
	tag, err := addTagFromDB(name)
	if err != nil {
		return err
	}
	if src.Setting.Redis {
		err := addTagFromCache(tag)
		if err != nil {
			return err
		}
	}
	return nil
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
			utils.ExceptionLog(err, msg)
			return err
		}
	}
	err := rc.Flush()
	utils.ExceptionLog(err, "flush fail when insert tag to cache")
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
			utils.ExceptionLog(err, msg)
			return nil, err
		}
	}
	utils.ExceptionLog(rc.Flush(), "Fail to flush query tag info")
	for i, field := range tagInfoCacheFields {
		result, err := rc.Receive()
		if err != nil {
			msg := fmt.Sprintf("Fail to do rc.Receive(), tagID = %v, field = %v", id, field)
			utils.ExceptionLog(err, msg)
			return nil, err
		}
		if result != nil {
			cacheReturnResults[i] = string(result.([]byte))
		} else {
			// 缓存失效
			msg := fmt.Sprintf("缓存未命中！tagID = %v, field = %v", id, field)
			utils.LogPlus(msg)
			tagFromDB, err := queryTagByIDFromDB(id)
			if err != nil {
				msg := fmt.Sprintf("Fail to query tag from DB when cache miss, tagID = %v", id)
				utils.ExceptionLog(err, msg)
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
	// Redis 中的时间改用时间戳存储
	uTime, _ := strconv.Atoi(cacheReturnResults[2])
	createTime := time.Unix(int64(uTime), 0)
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
	return queryTagByIDFromDB(id)
}
