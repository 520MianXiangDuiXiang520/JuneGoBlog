package message

import (
	junebao_top "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

type TalkingListResp struct {
	Header junebao_top.BaseRespHeader `json:"header"`
}

type TalkingListReq struct {
}

func (r TalkingListReq) JSON(ctx *gin.Context,
	jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}
