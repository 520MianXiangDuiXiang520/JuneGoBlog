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

func AuthRegister(rg *gin.RouterGroup) {
	rg.POST("/login", authLoginRoutes()...)
	rg.POST("/info", authInfoRoutes()...) // 获取用户信息
	rg.POST("/logout", authLogoutRoutes()...)
}

func authLoginRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneGin.EasyHandler(check.AuthLoginCheck,
			server.AuthLoginLogic, message.AuthLoginReq{}),
	}
}

func authInfoRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneGin.EasyHandler(check.AuthInfoCheck,
			server.AuthInfoLogic, message.AuthInfoReq{}),
	}
}

func authLogoutRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneGin.EasyHandler(check.AuthLogoutCheck,
			server.AuthLogoutLogic, message.AuthLogoutReq{}),
	}
}
