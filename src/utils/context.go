package utils

import (
	"JuneGoBlog/src/consts"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CheckFunc func(ctx *gin.Context, req interface{}) (interface{}, error)
type LogicFunc func(ctx *gin.Context, req interface{}) interface{}

func EasyHandler(cf CheckFunc, lf LogicFunc, req interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		var resp interface{}
		if err := context.BindJSON(req); err != nil {
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
