package check

import (
	"JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func TagListCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	return nil, nil
}

func TagAddCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	reqA := req.(*message.TagAddReq)
	if reqA.TagName == "" {
		log.Printf("Check Add Tag Error! name = [%v]\n", reqA.TagName)
		return junebao_top.ParamErrorRespHeader, errors.New("")
	}
	return nil, nil
}
