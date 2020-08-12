package message

import "JuneGoBlog/src/dao"

// api/tag/list 请求格式
type TagListReq struct {}

// 标签信息
type TagInfo struct {
	dao.Tag
	ArticleTotal int       `json:"article_total"`  // 文章数
}

// api/tag/list 响应格式
type TagListResp struct {
	Header BaseRespHeader   `json:"header"`        // 响应头
	Total int               `json:"total"`         // 标签总数
	Tags []TagInfo          `json:"tags"`          // 标签列表
}

type TagAddReq struct {
	TagName string          `form:"name" json:"name"`          // 标签名
}

type TagAddResp struct {
	Header BaseRespHeader   `json:"header"`
}