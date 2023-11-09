package server

import (
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/consts"
	"JuneGoBlog/internal/db/old"
	"JuneGoBlog/internal/util"
	"fmt"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthLoginLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.AuthLoginReq)
	resp := message.AuthLoginResp{}
	// 得到 username 和 password
	password := util.Sha256(request.Password)
	user, ok := old.GetUser(request.Username, password)
	if !ok {
		return juneGin.UnauthorizedRespHeader
	}
	token := util.GetHashWithTimeUUID(user.Username + user.Password)
	err := old.InsertUserToken(user, token, time.Now().Add(consts.ExpireDuration))
	if err != nil {
		msg := fmt.Sprintf("insert userToken fail, user id = %v, token = %v\n", user.ID, token)
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Token = token
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func AuthInfoLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	resp := message.AuthInfoResp{}
	u, ok := ctx.Get("user")
	if !ok {
		return juneGin.UnauthorizedRespHeader
	}
	user := u.(*old.User)
	resp.ID = user.ID
	resp.Username = user.Username
	resp.Permiter = user.Permiter
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func AuthLogoutLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	resp := message.AuthLogoutResp{}
	u, ok := ctx.Get("user")
	if !ok {
		return juneGin.UnauthorizedRespHeader
	}
	user := u.(*old.User)
	err := old.DeleteUserTokenByUID(user.ID)
	if err != nil {
		msg := fmt.Sprintf("logout fail(delete user token fail), uid = %v", user.ID)
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}
