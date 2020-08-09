package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", utils.EasyHandler(&utils.LogicContext{CheckFunc: check.CheckList, LogicFunc: server.ListLogic}))
	rg.POST("/add", utils.EasyHandler(&utils.LogicContext{CheckFunc: check.CheckList, LogicFunc: server.ListLogic}))
	rg.POST("/update", utils.EasyHandler(&utils.LogicContext{CheckFunc: check.CheckList, LogicFunc: server.ListLogic}))
	rg.POST("/delete", utils.EasyHandler(&utils.LogicContext{CheckFunc: check.CheckList, LogicFunc: server.ListLogic}))
	rg.POST("/detail", utils.EasyHandler(&utils.LogicContext{CheckFunc: check.CheckList, LogicFunc: server.ListLogic}))
}
