package routes

import (
	"JuneGoBlog/src/check"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/junebao.top/middleware"
	"JuneGoBlog/src/message"
	middleware2 "JuneGoBlog/src/middleware"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func AuthRegister(rg *gin.RouterGroup) {
	rg.POST("/login", authLoginRoutes()...)
	rg.POST("/info", authInfoRoutes()...) // 获取用户信息
	rg.POST("/logout", authLogoutRoutes()...)
}

func authLoginRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.AuthLoginCheck,
			server.AuthLoginLogic, message.AuthLoginReq{}),
	}
}

func authInfoRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		junebao_top.EasyHandler(check.AuthInfoCheck,
			server.AuthInfoLogic, message.AuthInfoReq{}),
	}
}

func authLogoutRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		junebao_top.EasyHandler(check.AuthLogoutCheck,
			server.AuthLogoutLogic, message.AuthLogoutReq{}),
	}
}
