package message

import (
	"JuneGoBlog/src/dao"
	junebao_top "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

type TalkingListResp struct {
	Header  junebao_top.BaseRespHeader `json:"header"`
	HasNext bool                       `json:"hasNext"`
	Talks   []dao.Talks                `json:"talks"`
}

type TalkingListReq struct {
	ArticleID    int `json:"articleID"`    // 文章ID（必须）
	PageSize     int `json:"pageSize"`     // 每页评论数（必须）
	Page         int `json:"page"`         // 页码（必须）
	ParentTalkID int `json:"parentTalkID"` // 父评论 ID（非必须）
}

func (r TalkingListReq) JSON(ctx *gin.Context,
	jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

type TalkingAddResp struct {
	Header junebao_top.BaseRespHeader `json:"header"`
}

type TalkingAddReq struct {
	ArticleID int    `json:"articleID"` // 文章ID（必须）
	Text      string `json:"text"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Type      int    `json:"type"`
	PTalkID   int    `json:"pTalkID"`
	SiteLink  string `json:"siteLink"`
}

func (r TalkingAddReq) JSON(ctx *gin.Context,
	jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}
