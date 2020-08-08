package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespHeader interface {}

type LogicContext struct {
	CheckFunc func(ctx *gin.Context) (RespHeader, error)
	LogicFunc func(ctx *gin.Context) RespHeader
}

func (lc *LogicContext) do() gin.HandlerFunc {
	 return func(context *gin.Context) {
		var resp interface{}
		if checkResp, err := lc.CheckFunc(context); err != nil {
			resp = checkResp
		} else {
			resp = lc.LogicFunc(context)
		}
		context.JSON(http.StatusOK, resp)
	}
}

func EasyHandler(lc *LogicContext) gin.HandlerFunc {
	return lc.do()
}
