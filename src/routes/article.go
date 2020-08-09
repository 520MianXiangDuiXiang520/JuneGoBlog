package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", utils.EasyHandler(check.ListCheck, server.ListLogic))
	rg.POST("/add", utils.EasyHandler(check.ListCheck, server.ListLogic))
	rg.POST("/update", utils.EasyHandler(check.ListCheck, server.ListLogic))
	rg.POST("/delete", utils.EasyHandler(check.ListCheck, server.ListLogic))
	rg.POST("/detail", utils.EasyHandler(check.ListCheck, server.ListLogic))
}
