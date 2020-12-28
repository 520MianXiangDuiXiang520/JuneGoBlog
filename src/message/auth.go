package message

import (
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
)

type AuthLoginResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
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
	Header   juneGin.BaseRespHeader `json:"header"`
	ID       int                    `json:"id"`
	Username string                 `json:"username"`
	Permiter int                    `json:"permiter"`
}

type AuthInfoReq struct {
}

func (r *AuthInfoReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}

type AuthLogoutResp struct {
	Header juneGin.BaseRespHeader `json:"header"`
}

type AuthLogoutReq struct {
}

func (r *AuthLogoutReq) JSON(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&r)
}
