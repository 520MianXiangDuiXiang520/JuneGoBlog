package junebao_top

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CheckFunc func(ctx *gin.Context, req BaseReqInter) (BaseRespInter, error)
type LogicFunc func(ctx *gin.Context, req BaseReqInter) BaseRespInter

// 解析请求，整合检查请求参数，响应逻辑，并响应
func EasyHandler(cf CheckFunc, lf LogicFunc, req BaseReqInter) gin.HandlerFunc {
	return func(context *gin.Context) {
		var resp interface{}
		if err := req.JSON(context, &req); err != nil {
			log.Printf("EasyHandler: BindJSON ERROR!!!")
			resp = ParamErrorRespHeader
		} else {
			if checkResp, err := cf(context, req); err != nil {
				resp = checkResp
			} else {
				resp = lf(context, req)
			}
		}
		context.Set("resp", resp)
		context.JSON(http.StatusOK, resp)
	}
}
