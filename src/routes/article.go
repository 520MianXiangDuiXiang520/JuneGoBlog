package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", utils.EasyHandler(check.ListCheck, server.ListLogic, &message.ArticleListReq{}))
	rg.POST("/add", utils.EasyHandler(check.ListCheck, server.ListLogic, message.ArticleListReq{}))
	rg.POST("/update", utils.EasyHandler(check.ListCheck, server.ListLogic, message.ArticleListReq{}))
	rg.POST("/delete", utils.EasyHandler(check.ListCheck, server.ListLogic, message.ArticleListReq{}))
	rg.POST("/detail", utils.EasyHandler(check.ListCheck, server.ListLogic, message.ArticleListReq{}))
}
