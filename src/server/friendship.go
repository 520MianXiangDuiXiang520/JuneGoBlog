package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"fmt"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/gin-gonic/gin"
)

// 友链审批
func FriendApprovalLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	reqA := req.(*message.FriendApprovalReq)
	resp := message.FriendApprovalResp{}
	if err := dao.UpdateFriendStatusByID(reqA.FriendshipID, reqA.Result); err != nil {
		msg := fmt.Sprintf("update friendship statue error! fid = [%v]", reqA.FriendshipID)
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

// 获取友链列表（展示）
func FriendShipListLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	resp := message.FriendShipListResp{}
	// 从数据库中读取所有的 friendShip Link 信息
	friendshipList := make([]dao.FriendShipLink, 0)
	if err := dao.QueryAllFriendLinkByStatus(consts.FriendShipApprovalPass, &friendshipList); err != nil {
		msg := fmt.Sprintf("FriendShipListLogic dao.DB.Find ERROR [%v]\n", err)
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}

// 获取不展示的友链列表
func FriendUnShowListLogic(ctx *gin.Context, re juneGin.BaseReqInter) juneGin.BaseRespInter {
	resp := message.FriendUnShowListResp{}
	reqU := re.(*message.FriendUnShowListReq)
	friendshipList := make([]dao.FriendShipLink, 0)
	if reqU.Status == 0 {
		hopeStatus := [2]int{consts.FriendShipApproving, consts.FriendShipApprovalFail}
		if err := dao.QueryAllFriendLinkINStatus(hopeStatus[:], &friendshipList); err != nil {
			msg := fmt.Sprintf("FriendUnShowListLogic dao.DB.Find ERROR [%v]\n", err)
			juneLog.ExceptionLog(err, msg)
			return juneGin.SystemErrorRespHeader
		}
	} else {
		if err := dao.QueryAllFriendLinkByStatus(reqU.Status, &friendshipList); err != nil {
			msg := fmt.Sprintf("FriendUnShowListLogic dao.DB.Find ERROR [%v]\n", err)
			juneLog.ExceptionLog(err, msg)
			return juneGin.SystemErrorRespHeader
		}
	}
	resp.Header = juneGin.SuccessRespHeader
	resp.FriendShipList = friendshipList
	resp.Total = len(friendshipList)
	return resp
}

// 申请添加友链
func FriendApplicationLogic(ctx *gin.Context, re juneGin.BaseReqInter) juneGin.BaseRespInter {
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
		msg := fmt.Sprintf("FriendApplicationLogic CALL AddFriendship Error !!!")
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}

	resp.Header = juneGin.SuccessRespHeader
	return resp
}

// 删除友链
func FriendDeleteLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	reqD := req.(*message.FriendDeleteReq)
	var resp message.FriendApplicationResp
	if err := dao.DeleteFriendshipByID(reqD.ID); err != nil {
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}
