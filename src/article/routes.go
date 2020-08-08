package article

import (
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup) {
	rg.POST("/list", utils.EasyHandler(&utils.LogicContext{CheckFunc: CheckList, LogicFunc: ListLogic}))
	rg.POST("/add", utils.EasyHandler(&utils.LogicContext{CheckFunc: CheckList, LogicFunc: ListLogic}))
	rg.POST("/update", utils.EasyHandler(&utils.LogicContext{CheckFunc: CheckList, LogicFunc: ListLogic}))
	rg.POST("/delete", utils.EasyHandler(&utils.LogicContext{CheckFunc: CheckList, LogicFunc: ListLogic}))
	rg.POST("/detail", utils.EasyHandler(&utils.LogicContext{CheckFunc: CheckList, LogicFunc: ListLogic}))
}