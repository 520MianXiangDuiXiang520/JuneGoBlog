package dao

// 单个标签信息
type TagInfo struct {
	Name string       `form:"name"`              // 标签名
	CreateTime int    `form:"createTime"`        // 创建时间（时间戳）
	ArticleAmount int `form:"articleAmount"`     // 该标签下的文章数
}

// 单个文章信息
type ArticleInfo struct {
	Name string       `json:"name"`              // 文章名
	CreateTime int    `json:"createTime"`        // 创建时间（时间戳）
	Abstract string   `json:"abstract"`          // 摘要
	Text string       `json:"text"`              // 正文
	ReadingAmount int `json:"readingAmount"`     // 阅读量
	Tags []TagInfo    `json:"tags"`              // 标签信息
}


// 友链信息
type FriendShipLink struct {
	ID int            `json:"id" gorm:"column:id"`
	SiteName string   `json:"siteName" gorm:"column:siteName"`       // 网站名
	SiteLink string   `json:"link" gorm:"column:siteLink"`           // 链接
	ImgLink string    `json:"imgLink" gorm:"column:imgLink"`         // 网站图标链接
	Intro string      `json:"intro" gorm:"column:intro"`             // 简介
}

func (FriendShipLink) TableName() string {
	return "friendship"
}