package dao

// 单个标签信息
type TagInfo struct {
	Name string         `form:"name"`            // 标签名
	CreateTime int      `form:"createTime"`      // 创建时间（时间戳）
	ArticleAmount int   `form:"articleAmount"`   // 该标签下的文章数
}

// 单个文章信息
type ArticleInfo struct {
	Name string         `json:"name"`            // 文章名
	CreateTime int      `json:"createTime"`      // 创建时间（时间戳）
	Abstract string     `json:"abstract"`        // 摘要
	Text string         `json:"text"`            // 正文
	ReadingAmount int   `json:"readingAmount"`   // 阅读量
	Tags []TagInfo      `json:"tags"`            // 标签信息
}
