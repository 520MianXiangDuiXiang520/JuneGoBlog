package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

func FriendShipListLogic(ctx *gin.Context, req interface{}) interface{} {
	resp := message.FriendShipListResp{}

	// 从数据库中读取所有的 friendShip Link 信息
	friendshipList := make([]dao.FriendShipLink, 0)
	if err := dao.QueryAllFriendLink(&friendshipList); err != nil {
		log.Printf("FriendShipListLogic dao.DB.Find ERROR [%v]\n", err)
		return consts.SystemErrorRespHeader
	}
	resp.Header = consts.SuccessRespHeader
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}

// 添加友链
func FriendAddLogic(ctx *gin.Context, re interface{}) interface{} {
	req := re.(*message.FriendAddReq)
	resp := message.FriendAddResp{}

	err := dao.AddFriendship(&dao.FriendShipLink{
		SiteName: req.SiteName,
		SiteLink: req.SiteLink,
		ImgLink:  req.ImgLink,
		Intro:    req.Intro,
	})

	if err != nil {
		log.Printf("FriendAddLogic CALL AddFriendship Error !!!")
		return consts.SystemErrorRespHeader
	}

	resp.Header = consts.SuccessRespHeader
	return resp
}
