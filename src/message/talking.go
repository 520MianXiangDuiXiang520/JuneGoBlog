package message

import (
	"JuneGoBlog/src/dao"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
)

type TalkingListResp struct {
	Header  juneGin.BaseRespHeader `json:"header"`
	HasNext bool                   `json:"hasNext"`
	Talks   []dao.Talks            `json:"talks"`
}

type TalkingListReq struct {
	ArticleID    int `json:"articleID" check:"not null"`         // 文章ID（必须）
	PageSize     int `json:"pageSize" check:"not null; more: 0"` // 每页评论数（必须）
	Page         int `json:"page" check:"not null; more: 0"`     // 页码（必须）
	ParentTalkID int `json:"parentTalkID"`                       // 父评论 ID（非必须）
}

func (r *TalkingListReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type TalkingAddResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}

type TalkingAddReq struct {
	ArticleID int    `json:"articleID" check:"not null"` // 文章ID（必须）
	Text      string `json:"text" check:"not null"`
	Username  string `json:"username"`
	Email     string `json:"email" check:"not null; email"`
	Type      int    `json:"type"`
	PTalkID   int    `json:"pTalkID"`
	SiteLink  string `json:"siteLink"`
}

func (r *TalkingAddReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}
