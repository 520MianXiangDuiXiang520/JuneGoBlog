package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/message"
	"errors"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
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

func AuthLoginCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.AuthLoginReq)
	if checkUsername(request.Username) && checkPassword(request.Password) {
		return http.StatusOK, nil
	}
	return juneGin.ParamErrorRespHeader, errors.New("")
}

func AuthInfoCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	return http.StatusOK, nil
}

func AuthLogoutCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	return http.StatusOK, nil
}
