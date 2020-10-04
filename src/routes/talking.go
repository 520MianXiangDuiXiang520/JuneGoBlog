package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"

	//"JuneGoBlog/src/junebao.top"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func TalkingRegister(rg *gin.RouterGroup) {
	rg.POST("/list", talkingListRoutes()...)
	rg.POST("/add", talkingAddRoutes()...)
}

func talkingListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.TalkingListCheck,
			server.TalkingListLogic, message.TalkingListReq{}),
	}
}

func talkingAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.TalkingAddCheck,
			server.TalkingAddLogic, message.TalkingAddReq{}),
	}
}
