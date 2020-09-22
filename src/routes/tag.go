package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/junebao.top"
	middleware2 "JuneGoBlog/src/junebao.top/middleware"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func TagRegister(rg *gin.RouterGroup) {
	rg.POST("/list", tagListRoutes()...)
	rg.POST("/add", tagAddRoutes()...)
	rg.POST("/delete", tagDeleteRoutes()...)
}

func tagDeleteRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.Permiter(middleware.AdminPermit),
		// TODO: 添加删除逻辑
	}
}

func tagAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.Permiter(middleware.AdminPermit),
		junebao_top.EasyHandler(check.TagAddCheck, server.TagAddLogin, &message.TagAddReq{}),
	}
}

func tagListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.TagListCheck, server.TagListLogin, &message.TagListReq{}),
	}
}
