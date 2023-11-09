package routes

import (
	"JuneGoBlog/internal/api/check"
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/api/server"
	middleware2 "JuneGoBlog/internal/middleware"
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
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneMiddleware.Permiter(middleware2.AdminPermit),
		// TODO: 添加删除逻辑
	}
}

func tagAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneMiddleware.Permiter(middleware2.AdminPermit),
		juneGin.EasyHandler(check.TagAddCheck, server.TagAddLogin, message.TagAddReq{}),
	}
}

func tagListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.TagListCheck, server.TagListLogin, message.TagListReq{}),
	}
}
