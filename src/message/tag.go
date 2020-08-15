package message

import (
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

// api/tag/list 请求格式
type TagListReq struct{}

func (tlq TagListReq) JSON(ctx *gin.Context, jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

// 标签信息
type TagInfo struct {
	dao.Tag
	ArticleTotal int `json:"article_total"` // 文章数
}

// api/tag/list 响应格式
type TagListResp struct {
	Header junebao_top.BaseRespHeader `json:"header"` // 响应头
	Total  int                        `json:"total"`  // 标签总数
	Tags   []TagInfo                  `json:"tags"`   // 标签列表
}

type TagAddReq struct {
	TagName string `form:"name" json:"name"` // 标签名
}

func (taq TagAddReq) JSON(ctx *gin.Context, jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type TagAddResp struct {
	Header junebao_top.BaseRespHeader `json:"header"`
}
