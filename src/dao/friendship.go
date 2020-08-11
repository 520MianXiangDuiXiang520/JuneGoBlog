package dao

import "log"

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