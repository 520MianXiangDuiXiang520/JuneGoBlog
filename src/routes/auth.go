package routes

import (
	"JuneGoBlog/src/check"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func AuthRegister(rg *gin.RouterGroup) {
	rg.POST("/login", authLoginRoutes()...)

}

func authLoginRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.AuthLoginCheck,
			server.AuthLoginLogic, &message.AuthLoginReq{}),
	}
}
