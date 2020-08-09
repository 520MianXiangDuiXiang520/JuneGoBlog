package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func FriendShipListLogic(ctx *gin.Context) utils.RespHeader {
	resp := message.FriendShipListResp{}

	// 从数据库中读取所有的 friendShip Link 信息
	friendshipList := make([]dao.FriendShipLink, 0)
	if err := dao.QueryAllFriendLink(&friendshipList); err != nil {
		log.Printf("FriendShipListLogic dao.DB.Find ERROR [%v]\n", err)
		return consts.SystemError
	}
	resp.Header = consts.SuccessHead
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}
