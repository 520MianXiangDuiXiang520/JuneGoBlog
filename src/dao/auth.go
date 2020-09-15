package dao

import (
	"time"
)

/**
* WiKi: 通过用户名和密码查询用户，如果用户不存在第二个参数返回 false
* Author: JuneBao
* Time: 2020/9/15 20:17
**/
func GetUser(username, password string) (*User, bool) {
	var user User
	err := DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	return &user, err == nil
}

func GetUserByToken(token string) (*User, bool) {
	var ut UserToken
	var user User
	err := DB.Where("token = ?", token).First(&ut).Error
	if err != nil {
		return nil, false
	}
	if ut.ExpireTime.Unix() < time.Now().Unix() {
		err = DeleteUserTokenByID(ut.ID)
		return nil, false
	}
	err = DB.Where("id = ?", ut.UserID).First(&user).Error
	return &user, err == nil
}

func DeleteUserTokenByUID(uid int) error {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Where("user_id = ?", uid).Delete(&UserToken{}).Error
	return err
}

func DeleteUserTokenByID(id int) error {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = tx.Where("id = ?", id).Delete(&UserToken{}).Error
	return err
}

func InsertUserToken(user *User, token string, expire time.Time) error {
	tx := DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	err = DeleteUserTokenByUID(user.ID)
	if err != nil {
		return err
	}
	err = tx.Create(&UserToken{
		Token:      token,
		UserID:     user.ID,
		ExpireTime: expire,
	}).Error
	return err
}
