package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func FriendShipRoutes (rg *gin.RouterGroup) {
	rg.POST("list/", utils.EasyHandler(check.FriendShipListCheck, server.FriendShipListLogic))
}
