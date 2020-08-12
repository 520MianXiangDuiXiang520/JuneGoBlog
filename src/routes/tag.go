package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func TagRegister(rg *gin.RouterGroup) {
	rg.POST("/list", utils.EasyHandler(check.TagListCheck,
		server.TagListLogin, &message.TagListReq{}))
	rg.POST("/add", utils.EasyHandler(check.TagAddCheck,
		server.TagAddLogin,  &message.TagAddReq{}))
	rg.POST("/delete")
}
