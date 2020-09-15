package check

import (
	"JuneGoBlog/src/consts"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkUsername(username string) bool {
	return len(username) >= consts.MinUsernameLength &&
		len(username) <= consts.MaxUsernameLength
}

func checkPassword(password string) bool {
	return len(password) > 0
}

func AuthLoginCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	request := req.(*message.AuthLoginReq)
	if checkUsername(request.Username) && checkPassword(request.Password) {
		return http.StatusOK, nil
	}
	return junebao_top.ParamErrorRespHeader, errors.New("")
}
