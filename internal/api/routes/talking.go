package routes

import (
	"JuneGoBlog/internal/api/check"
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/api/server"
	"JuneGoBlog/internal/middleware"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
)

func TalkingRegister(rg *gin.RouterGroup) {
	rg.POST("/list", talkingListRoutes()...)
	rg.POST("/add", talkingAddRoutes()...)
}

func talkingListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.TalkingListCheck,
			server.TalkingListLogic, message.TalkingListReq{}),
	}
}

func talkingAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneGin.EasyHandler(check.TalkingAddCheck,
			server.TalkingAddLogic, message.TalkingAddReq{}),
	}
}
