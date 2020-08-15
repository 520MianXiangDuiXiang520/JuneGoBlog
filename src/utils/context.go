package utils

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CheckFunc func(ctx *gin.Context, req message.BaseReqInter) (message.BaseRespInter, error)
type LogicFunc func(ctx *gin.Context, req message.BaseReqInter) message.BaseRespInter

func EasyHandler(cf CheckFunc, lf LogicFunc, req message.BaseReqInter) gin.HandlerFunc {
	return func(context *gin.Context) {
		var resp interface{}
		if err := req.JSON(context, &req); err != nil {
			log.Printf("EasyHandler: BindJSON ERROR!!!")
			resp = consts.ParamErrorRespHeader
		} else {
			if checkResp, err := cf(context, req); err != nil {
				resp = checkResp
			} else {
				resp = lf(context, req)
			}
		}
		context.JSON(http.StatusOK, resp)
	}
}
