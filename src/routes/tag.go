package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/server"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneMiddleware "github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	"github.com/gin-gonic/gin"
)

func TagRegister(rg *gin.RouterGroup) {
	rg.POST("/list", tagListRoutes()...)
	rg.POST("/add", tagAddRoutes()...)
	rg.POST("/delete", tagDeleteRoutes()...)
}

func tagDeleteRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		// TODO: 添加删除逻辑
	}
}

func tagAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.TagAddCheck, server.TagAddLogin, message.TagAddReq{}),
	}
}

func tagListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.TagListCheck, server.TagListLogin, message.TagListReq{}),
	}
}
