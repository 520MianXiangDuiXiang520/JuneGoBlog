package server

import (
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

func TalkingListLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.TalkingListReq)
	resp := message.TalkingListResp{}
	// TODO:...
	log.Println(request)
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}
