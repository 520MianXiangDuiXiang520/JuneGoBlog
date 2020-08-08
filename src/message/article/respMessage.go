package article

import (
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
)



type ListResp struct {
	Header message.BaseRespHeader    `json:"header"`        // 响应头
	ArticleList []dao.ArticleInfo    `json:"articleList"`   // 文章列表
	Total int                        `json:"total"`         // 将返回的文章总数
}
