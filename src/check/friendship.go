package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FriendShipListCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	// 无请求参数，不需要校验
	return http.StatusOK, nil
}

func FriendShipUnShowListCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	reqU := req.(*message.FriendUnShowListReq)
	hopeStatus := [2]int{consts.FriendShipApproving, consts.FriendShipApprovalFail}
	if reqU.Status != 0 {
		for _, hs := range hopeStatus {
			if reqU.Status == hs {
				return nil, nil
			}
		}
		return junebao_top.ParamErrorRespHeader, errors.New("")
	}
	return nil, nil
}

func FriendApprovalCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	reqA := req.(*message.FriendApprovalReq)

	// 检查 FriendshipID 是否存在
	if _, ok := dao.HasFriendLinkByID(reqA.FriendshipID); !ok {
		return junebao_top.ParamErrorRespHeader, errors.New("NO Result")
	}
	// 检查 result
	TrueStatus := [2]int{
		consts.FriendShipApprovalPass,
		consts.FriendShipApprovalFail,
	}
	for _, s := range TrueStatus {
		if s == reqA.Result {
			return nil, nil
		}
	}
	return junebao_top.ParamErrorRespHeader, errors.New("BadParam")
}

func FriendApplicationCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	reqF := req.(*message.FriendApplicationReq)
	// name 和 Link 必填
	if reqF.SiteName == "" || reqF.SiteLink == "" {
		return junebao_top.ParamErrorRespHeader, errors.New("参数异常")
	}
	return nil, nil
}

func FriendDeleteCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	reqD := req.(*message.FriendDeleteReq)
	if reqD.ID <= 0 {
		return junebao_top.ParamErrorRespHeader, errors.New("请求参数错误")
	}
	return nil, nil
}
