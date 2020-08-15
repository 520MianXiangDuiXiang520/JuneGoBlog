package message

import "github.com/gin-gonic/gin"

type BaseRespHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BaseReqInter interface {
	JSON(ctx *gin.Context, jsonReq *BaseReqInter) error
}

type BaseRespInter = interface{}
