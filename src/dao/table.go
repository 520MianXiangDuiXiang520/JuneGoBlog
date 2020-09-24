package dao

import "time"

// 友链信息
type FriendShipLink struct {
	ID       int    `json:"id" gorm:"column:id"`
	SiteName string `json:"siteName" gorm:"column:siteName"` // 网站名
	SiteLink string `json:"link" gorm:"column:siteLink"`     // 链接
	ImgLink  string `json:"imgLink" gorm:"column:imgLink"`   // 网站图标链接
	Intro    string `json:"intro" gorm:"column:intro"`       // 简介
	Status   int    `json:"status" gorm:"column:status"`     // 状态
}

func (FriendShipLink) TableName() string {
	return "friendship"
}

// 文章标签
type Tag struct {
	ID         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (Tag) TableName() string {
	return "tags"
}

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Permiter int    `json:"permiter" gorm:"column:permit"`
}

func (u User) GetID() int {
	return u.ID
}

func (User) TableName() string {
	return "users"
}

type UserToken struct {
	ID         int       `json:"id" gorm:"column:id"`
	UserID     int       `json:"userId" gorm:"column:user_id"`
	Token      string    `json:"token" gorm:"column:token"`
	ExpireTime time.Time `json:"createTime" gorm:"expire_time"`
}

func (UserToken) TableName() string {
	return "user_token"
}

type Article struct {
	ID         int       `json:"id" gorm:"column:id"`
	Title      string    `json:"title" gorm:"column:title"`
	Abstract   string    `json:"abstract" gorm:"column:abstract"`
	Text       string    `json:"text" gorm:"column:text"`
	AuthorID   int       `json:"authorID" gorm:"author_id"`
	CreateTime time.Time `json:"createTime" gorm:"create_time"`
}

// 文章信息
type ArticleListInfo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	Author     string `json:"author"`
	CreateTime int64  `json:"createTime"`
	Tags       []Tag  `json:"tags"`
}

func (Article) TableName() string {
	return "articles"
}

type ArticleTags struct {
	ID        int `json:"id"gorm:"column:id"`
	ArticleID int `json:"articleID"gorm:"column:article_id"`
	TagID     int `json:"tagID"gorm:"column:tag_id"`
}

func (ArticleTags) TableName() string {
	return "article_tags"
}
