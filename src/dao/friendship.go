package dao

import (
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	"log"
)

// 查询所有的友链信息
func QueryAllFriendLink(fl *[]FriendShipLink) error {
	return juneDao.GetDB().Find(&fl).Error
}

// 查询所有状态为 status 的友链
func QueryAllFriendLinkByStatus(status int, fl *[]FriendShipLink) error {
	return juneDao.GetDB().Where("status = ?", status).Find(&fl).Error
}

// 查询所有状态 in (status) 的友链
func QueryAllFriendLinkINStatus(status []int, fl *[]FriendShipLink) error {
	return juneDao.GetDB().Where("status IN (?)", status).Find(&fl).Error
}

// 根据 FID 判断某条友链是否存在
func HasFriendLinkByID(fid int) (*FriendShipLink, bool) {
	fl := new(FriendShipLink)
	juneDao.GetDB().Where("id = ?", fid).Find(&fl)
	if fl.Status == 0 && fl.ID == 0 {
		return nil, false
	} else {
		return fl, true
	}
}

func UpdateFriendStatusByID(fid, status int) error {
	tx := juneDao.GetDB().Begin()
	if err := tx.Error; err != nil {
		log.Printf("UpdateFriendStatusByID Begin Error, fid = [%v], : [%v]\n", fid, err)
		return err
	}
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("UpdateFriendStatusByID Error, Callbacked... fid = [%v]; [%v]", fid, err)
		}
		tx.Commit()
	}()
	return tx.Model(&FriendShipLink{ID: fid}).Update("status", status).Error
}

// 添加一条友链
func AddFriendship(fs *FriendShipLink) error {
	tx := juneDao.GetDB().Begin()
	if err := tx.Error; err != nil {
		log.Printf("AddFriendship Begin Error, FriendShipLink.Name = [%v], : [%v]\n", fs.SiteName, err)
		return err
	}
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("AddFriendship Error, Callbacked... FriendShipLink.Name = [%v]; [%v]", fs.SiteName, err)
		}
		tx.Commit()
	}()
	err = tx.Create(&fs).Error
	return err
}

func DeleteFriendshipByID(fid int) error {
	tx := juneDao.GetDB().Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("DeleteFriendshipByID DB.Begin() Error, ID is [%v]\n\n", err)
		}
		tx.Commit()
	}()
	return tx.Where("id = ?", fid).Delete(&FriendShipLink{}).Error
}
