package server

import (
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthLoginLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.AuthLoginReq)
	resp := message.AuthLoginResp{}
	// 得到 username 和 password
	password := util.Sha256(request.Password)
	user, ok := dao.GetUser(request.Username, password)
	if !ok {
		return junebaotop.UnauthorizedRespHeader
	}
	token := util.GetHashWithTimeUUID(user.Username + user.Password)
	expire, _ := time.ParseDuration("30m")
	err := dao.InsertUserToken(user, token, time.Now().Add(expire))
	if err != nil {
		msg := fmt.Sprintf("insert userToken fail, user id = %v, token = %v\n", user.ID, token)
		util.ExceptionLog(err, msg)
		return junebaotop.SystemErrorRespHeader
	}
	resp.Token = token
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}
