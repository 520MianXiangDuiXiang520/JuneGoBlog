package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func TagListCheck(ctx *gin.Context, req message.BaseReqInter) (message.BaseRespInter, error) {
	return nil, nil
}

func TagAddCheck(ctx *gin.Context, req message.BaseReqInter) (message.BaseRespInter, error) {
	reqA := req.(*message.TagAddReq)
	if reqA.TagName == "" {
		log.Printf("Check Add Tag Error! name = [%v]\n", reqA.TagName)
		return consts.ParamErrorRespHeader, errors.New("")
	}
	return nil, nil
}
