package message

import (
	"JuneGoBlog/internal/db/old"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
)

type ArticleTagsResp struct {
	Header juneGin.BaseRespHeader `json:"header"` // 响应头
	ID     int                    `json:"id"`
	Tags   []TagInfo              `json:"tags"`
}

type ArticleTagsReq struct {
	ArticleID int `json:"articleID" form:"articleID" check:"more: 0"`
}

type ArticleListResp struct {
	Header      juneGin.BaseRespHeader `json:"header"`      // 响应头
	ArticleList []old.ArticleInfo      `json:"articleList"` // 文章列表
	Total       int                    `json:"total"`       // 将返回的文章总数
}

// 请求文章列表格式
type ArticleListReq struct {
	Page     int `json:"page" form:"page" check:"not null; more: 0"`              // 页数
	PageSize int `json:"pageSize" form:"pageSize" check:"not null; size: [5,20]"` // 每页请求的文章数量
	Tag      int `json:"tag" form:"tag"`                                          // 标签
}

type ArticleDetailReq struct {
	ArticleID int `json:"articleId" check:"more: 0"`
}

type ArticleDetailResp struct {
	juneGin.BaseRespHeader
	old.Article
	Text string `json:"text"`
}

func (atr *ArticleTagsReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&atr)
}

func (flr *ArticleListReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&flr)
}

func (adr *ArticleDetailReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&adr)
}

type ArticleAddResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}

type ArticleAddReq struct {
	Title    string `json:"title" check:"not null"`
	Abstract string `json:"abstract"`
	Text     string `json:"text" check:"not null"`
	Tags     []int  `json:"tags"`
}

func (r *ArticleAddReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type ArticleUpdateResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}

type ArticleUpdateReq struct {
	ID       int    `json:"id" check:"not null"`
	Title    string `json:"title" check:"not null"`
	Abstract string `json:"abstract"`
	Text     string `json:"text" check:"not null"`
	Tags     []int  `json:"tags"`
}

func (r *ArticleUpdateReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type ArticleDeleteResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}

type ArticleDeleteReq struct {
	ID int `json:"id" check:"not null"`
}

func (r *ArticleDeleteReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}
