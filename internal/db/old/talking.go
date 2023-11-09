package old

import (
	"JuneGoBlog/internal"
	"fmt"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
)

func hasTalkWithDB(id int) bool {
	talk := &Talks{}
	juneDao.GetDB().Where("id = ?", id).First(talk)
	return talk.ID != 0
}

func HasTalk(id int) bool {
	if internal.GetSetting().Others.Redis {
		// TODO: BitMap
	}
	return hasTalkWithDB(id)
}

func addTalkWithDB(talk *Talks) error {
	tx := juneDao.GetDB().Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			msg := fmt.Sprintf("Failed to add comment, rolled back, talk = %v", talk)
			juneLog.ExceptionLog(err, msg)
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
	err := juneDao.GetDB().Where("article_id = ?", articleID).Offset(offset).Limit(pageSize).Find(&talks).Error
	return talks, err
}

func QueryTalkByTalkID(talkID int) (res Talks, err error) {
	err = juneDao.GetDB().Where("id = ?", talkID).First(&res).Error
	return res, err
}
