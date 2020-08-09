package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespHeader interface{}

type CheckFunc func(ctx *gin.Context) (RespHeader, error)
type LogicFunc func(ctx *gin.Context) RespHeader

func EasyHandler(cf CheckFunc, lf LogicFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		var resp interface{}
		if checkResp, err := cf(context); err != nil {
			resp = checkResp
		} else {
			resp = lf(context)
		}
		context.JSON(http.StatusOK, resp)
	}
}
