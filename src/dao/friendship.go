package dao

import (

	"log"
)

// 查询所有的友链信息
func QueryAllFriendLink(fl *[]FriendShipLink) error {
	return DB.Find(&fl).Error
}

// 添加一条友链
func AddFriendship(fs *FriendShipLink) error {
	tx := DB.Begin()
	if err := tx.Error; err != nil{
		log.Printf("AddFriendship Begin Error, FriendShipLink.Name = [%v], : [%v]\n",fs.SiteName, err)
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
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("DeleteFriendshipByID DB.Begin() Error, ID is [%v]\n\n", err)
		}
		tx.Commit()
	}()

	deleteFri := tx.Model(&FriendShipLink{}).Where("id = ?", fid)
	var count int
	deleteFri.Count(&count)
	if count <= 0 {
		log.Printf("No This Friendship")
		return NoRecordError
	}
	return tx.Where("id = ?", fid).Delete(&FriendShipLink{}).Error
}