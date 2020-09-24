package message

import (
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

type ArticleTagsResp struct {
	Header junebaotop.BaseRespHeader `json:"header"` // 响应头
	ID     int                       `json:"id"`
	Tags   []TagInfo                 `json:"tags"`
}

type ArticleTagsReq struct {
	ArticleID int `json:"articleID"form:"articleID"`
}

type ArticleListResp struct {
	Header      junebaotop.BaseRespHeader `json:"header"`      // 响应头
	ArticleList []dao.ArticleListInfo     `json:"articleList"` // 文章列表
	Total       int                       `json:"total"`       // 将返回的文章总数
}

// 请求文章列表格式
type ArticleListReq struct {
	Page     int `json:"page"form:"page"`         // 页数
	PageSize int `json:"pageSize"form:"pageSize"` // 每页请求的文章数量
	Tag      int `json:"tag"form:"tag"`           // 标签
}

type ArticleDetailReq struct {
	ArticleID int `json:"articleId"`
}

type ArticleDetailResp struct {
	junebaotop.BaseRespHeader
	dao.Article
	Text string `json:"text"`
}

func (atr ArticleTagsReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

func (flr ArticleListReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

func (adr ArticleDetailReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type ArticleAddResp struct {
	Header junebaotop.BaseRespHeader `json:"header"`
}

type ArticleAddReq struct {
	dao.Article
	Tags []int `json:"tags"`
}

func (r ArticleAddReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type ArticleUpdateResp struct {
	Header junebaotop.BaseRespHeader `json:"header"`
}

type ArticleUpdateReq struct {
	dao.Article
	Tags []int `json:"tags"`
}

func (r ArticleUpdateReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type ArticleDeleteResp struct {
	Header junebaotop.BaseRespHeader `json:"header"`
}

type ArticleDeleteReq struct {
	ID int `json:"id"`
}

func (r ArticleDeleteReq) JSON(ctx *gin.Context,
	jsonReq *junebaotop.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}
