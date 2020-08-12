package dao

import "time"

// 友链信息
type FriendShipLink struct {
	ID int                `json:"id" gorm:"column:id"`
	SiteName string       `json:"siteName" gorm:"column:siteName"`       // 网站名
	SiteLink string       `json:"link" gorm:"column:siteLink"`           // 链接
	ImgLink string        `json:"imgLink" gorm:"column:imgLink"`         // 网站图标链接
	Intro string          `json:"intro" gorm:"column:intro"`             // 简介
}

func (FriendShipLink) TableName() string {
	return "friendship"
}

// 文章标签
type Tag struct {
	ID int                `json:"id" gorm:"column:id"`
	Name string           `json:"name" gorm:"column:name"`
	CreateTime time.Time  `json:"create_time" gorm:"column:create_time"`
}

func (Tag) TableName() string {
	return "tags"
}