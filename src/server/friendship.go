package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

// 友链审批
func FriendApprovalLogic(ctx *gin.Context, req junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	reqA := req.(*message.FriendApprovalReq)
	resp := message.FriendApprovalResp{}
	if err := dao.UpdateFriendStatusByID(reqA.FriendshipID, reqA.Result); err != nil {
		log.Printf("update friendship statue error! fid = [%v]", reqA.FriendshipID)
		return junebao_top.SystemErrorRespHeader
	}
	resp.Header = junebao_top.SuccessRespHeader
	return resp
}

// 获取友链列表（展示）
func FriendShipListLogic(ctx *gin.Context, req junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	resp := message.FriendShipListResp{}

	// 从数据库中读取所有的 friendShip Link 信息
	friendshipList := make([]dao.FriendShipLink, 0)
	if err := dao.QueryAllFriendLinkByStatus(consts.FriendShipApprovalPass, &friendshipList); err != nil {
		log.Printf("FriendShipListLogic dao.DB.Find ERROR [%v]\n", err)
		return junebao_top.SystemErrorRespHeader
	}
	resp.Header = junebao_top.SuccessRespHeader
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}

// 获取不展示的友链列表
func FriendUnShowListLogic(ctx *gin.Context, re junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	resp := message.FriendUnShowListResp{}
	reqU := re.(*message.FriendUnShowListReq)
	friendshipList := make([]dao.FriendShipLink, 0)
	if reqU.Status == 0 {
		hopeStatus := [2]int{consts.FriendShipApproving, consts.FriendShipApprovalFail}
		if err := dao.QueryAllFriendLinkINStatus(hopeStatus[:], &friendshipList); err != nil {
			log.Printf("FriendUnShowListLogic dao.DB.Find ERROR [%v]\n", err)
			return junebao_top.SystemErrorRespHeader
		}
	} else {
		if err := dao.QueryAllFriendLinkByStatus(reqU.Status, &friendshipList); err != nil {
			log.Printf("FriendUnShowListLogic dao.DB.Find ERROR [%v]\n", err)
			return junebao_top.SystemErrorRespHeader
		}
	}
	resp.Header = junebao_top.SuccessRespHeader
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}

// 申请添加友链
func FriendApplicationLogic(ctx *gin.Context, re junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	req := re.(*message.FriendApplicationReq)
	resp := message.FriendApplicationResp{}

	err := dao.AddFriendship(&dao.FriendShipLink{
		SiteName: req.SiteName,
		SiteLink: req.SiteLink,
		ImgLink:  req.ImgLink,
		Intro:    req.Intro,
		Status:   consts.FriendShipApproving,
	})

	if err != nil {
		log.Printf("FriendApplicationLogic CALL AddFriendship Error !!!")
		return junebao_top.SystemErrorRespHeader
	}

	resp.Header = junebao_top.SuccessRespHeader
	return resp
}

// 删除友链
func FriendDeleteLogic(ctx *gin.Context, req junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	reqD := req.(*message.FriendDeleteReq)
	var resp message.FriendApplicationResp
	if err := dao.DeleteFriendshipByID(reqD.ID); err != nil {
		return junebao_top.SystemErrorRespHeader
	}
	resp.Header = junebao_top.SuccessRespHeader
	return resp
}
