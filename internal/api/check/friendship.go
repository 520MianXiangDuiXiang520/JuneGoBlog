package check

import (
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/consts"
	"JuneGoBlog/internal/db/old"
	"errors"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FriendShipListCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	// 无请求参数，不需要校验
	return http.StatusOK, nil
}

func FriendShipUnShowListCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqU := req.(*message.FriendUnShowListReq)
	hopeStatus := [2]int{consts.FriendShipApproving, consts.FriendShipApprovalFail}
	if reqU.Status != 0 {
		for _, hs := range hopeStatus {
			if reqU.Status == hs {
				return nil, nil
			}
		}
		return juneGin.ParamErrorRespHeader, errors.New("")
	}
	return nil, nil
}

func FriendApprovalCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqA := req.(*message.FriendApprovalReq)

	// 检查 FriendshipID 是否存在
	if _, ok := old.HasFriendLinkByID(reqA.FriendshipID); !ok {
		return juneGin.ParamErrorRespHeader, errors.New("NO Result")
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
	return juneGin.ParamErrorRespHeader, errors.New("BadParam")
}

func FriendApplicationCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqF := req.(*message.FriendApplicationReq)
	// name 和 Link 必填
	if reqF.SiteName == "" || reqF.SiteLink == "" {
		return juneGin.ParamErrorRespHeader, errors.New("参数异常")
	}
	return nil, nil
}

func FriendDeleteCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqD := req.(*message.FriendDeleteReq)

	if _, ok := old.HasFriendLinkByID(reqD.ID); !ok {
		return juneGin.ParamErrorRespHeader, errors.New("请求参数错误")
	}
	return nil, nil
}
