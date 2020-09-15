package message

import (
	junebao_top "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
)

type AuthLoginResp struct {
	Header junebao_top.BaseRespHeader `json:"header"`
	Token  string
}

type AuthLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r AuthLoginReq) JSON(ctx *gin.Context,
	jsonReq *junebao_top.BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}
