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

func AuthRegister(rg *gin.RouterGroup) {
	rg.POST("/login", authLoginRoutes()...)
	rg.POST("/info", authInfoRoutes()...) // 获取用户信息
	rg.POST("/logout", authLogoutRoutes()...)
}

func authLoginRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.AuthLoginCheck,
			server.AuthLoginLogic, message.AuthLoginReq{}),
	}
}

func authInfoRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneMiddleware.Auth(middleware.TokenAuth),
		juneGin.EasyHandler(check.AuthInfoCheck,
			server.AuthInfoLogic, message.AuthInfoReq{}),
	}
}

func authLogoutRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneMiddleware.Auth(middleware.TokenAuth),
		juneGin.EasyHandler(check.AuthLogoutCheck,
			server.AuthLogoutLogic, message.AuthLogoutReq{}),
	}
}
