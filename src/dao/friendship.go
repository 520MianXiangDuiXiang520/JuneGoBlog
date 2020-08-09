package dao

// 查询所有的友链信息
func QueryAllFriendLink(fl *[]FriendShipLink) error {
	return DB.Find(&fl).Error
}
