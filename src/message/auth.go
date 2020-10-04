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

func (r *AuthLoginReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type AuthInfoResp struct {
	Header   junebao_top.BaseRespHeader `json:"header"`
	ID       int                        `json:"id"`
	Username string                     `json:"username"`
	Permiter int                        `json:"permiter"`
}

type AuthInfoReq struct {
}

func (r *AuthInfoReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type AuthLogoutResp struct {
	Header junebao_top.BaseRespHeader `json:"header"`
}

type AuthLogoutReq struct {
}

func (r *AuthLogoutReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}
