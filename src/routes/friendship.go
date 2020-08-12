package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func FriendShipRoutes (rg *gin.RouterGroup) {
	rg.POST("list/", utils.EasyHandler(check.FriendShipListCheck,
		server.FriendShipListLogic, &message.FriendAddReq{}))
	rg.POST("add/", utils.EasyHandler(check.FriendAddCheck,
		server.FriendAddLogic, &message.FriendAddReq{}))
	rg.POST("delete/", utils.EasyHandler(check.FriendShipListCheck,
		server.FriendDeleteLogic, &message.FriendDeleteReq{}))
	}
