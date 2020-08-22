package message

import (
	junebaotop "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

type ArticleTagInfo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	AuthorID   int    `json:"authorID"`
	CreateTime int64  `json:"createTime"`
	Tags       []TagInfo
}

type ArticleListResp struct {
	Header      junebaotop.BaseRespHeader `json:"header"`      // 响应头
	ArticleList []ArticleTagInfo          `json:"articleList"` // 文章列表
	Total       int                       `json:"total"`       // 将返回的文章总数
}

// 请求文章列表格式
type ArticleListReq struct {
	Page     int    `json:"page"form:"page"`         // 页数
	PageSize int    `json:"pageSize"form:"pageSize"` // 每页请求的文章数量
	Tag      string `json:"tag"form:"tag"`           // 标签
}

func (flr ArticleListReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

// 文章详情页请求
type ArticleDetailReq struct {
	ArticleID int `json:"articleId"`
}

func (adr ArticleDetailReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type ArticleDetailResp struct {
	junebaotop.BaseRespHeader
	ArticleTagInfo
	Text string `json:"text"`
}
