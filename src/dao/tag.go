package dao

import (
	"log"
	"time"
)

// 查询所有标签，按创建时间排序
func QueryAllTagsOrderByTime(resp *[]Tag) error {
	return DB.Order("create_time").Find(&resp).Error
}

func HasTagByName(name string) (*Tag, bool) {
	tag := new(Tag)
	DB.Where("name = ?", name).First(&tag)
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

func QueryTagByID(id int) Tag {
	result := Tag{}
	DB.Where("id = ?", id).First(&result)
	return result
}
