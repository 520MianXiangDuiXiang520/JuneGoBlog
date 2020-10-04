package dao

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/junebao.top/utils"
	"fmt"
)

func hasTalkWithDB(id int) bool {
	talk := &Talks{}
	DB.Where("id = ?", id).First(talk)
	return talk.ID != 0
}

func HasTalk(id int) bool {
	if src.Setting.Redis {
		// TODO: BitMap
	}
	return hasTalkWithDB(id)
}

func addTalkWithDB(talk *Talks) error {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			msg := fmt.Sprintf("Failed to add comment, rolled back, talk = %v", talk)
			utils.ExceptionLog(err, msg)
		}
		tx.Commit()
	}()
	err = tx.Create(talk).Error
	return err
}

func AddTalk(talk *Talks) error {
	return addTalkWithDB(talk)
}

func QueryTalksByArticleIDLimit(articleID, page, pageSize int) ([]Talks, error) {
	talks := make([]Talks, 0)
	offset := (page - 1) * pageSize
	err := DB.Where("article_id = ?", articleID).Offset(offset).Limit(pageSize).Find(&talks).Error
	return talks, err
}
