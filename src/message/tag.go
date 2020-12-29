package message

import (
	"JuneGoBlog/src/dao"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
)

// api/tag/list 请求格式
type TagListReq struct{}

func (tlq *TagListReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&tlq)
}

// api/tag/list 响应格式
type TagListResp struct {
	Header juneGin.BaseRespHeader `json:"header"` // 响应头
	Total  int                    `json:"total"`  // 标签总数
	Tags   []dao.Tag              `json:"tags"`   // 标签列表
}

type TagAddReq struct {
	TagName string `form:"name" json:"name"` // 标签名
}

func (taq *TagAddReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&taq)
}

type TagAddResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}
