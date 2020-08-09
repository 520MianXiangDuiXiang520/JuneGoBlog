package check

import (
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckList(ctx *gin.Context) (utils.RespHeader, error) {
	req := new(message.ListReq)
	var resp *message.BaseRespHeader
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Println("Bind JSON Error!!!")
		resp = &message.BaseRespHeader{Code: http.StatusBadRequest, Msg: "请求参数错误"}
		return resp, errors.New("")
	}
	return http.StatusOK, nil
}
