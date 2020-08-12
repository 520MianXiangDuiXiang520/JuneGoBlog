package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FriendShipListCheck(ctx *gin.Context, req interface{}) (interface{}, error) {
	// 无请求参数，不需要校验
	return http.StatusOK, nil
}

func FriendAddCheck(ctx *gin.Context, req interface{}) (interface{}, error) {
	reqF := req.(*message.FriendAddReq)
	// name 和 Link 必填
	if reqF.SiteName == "" || reqF.SiteLink == "" {
		return consts.ParamErrorRespHeader, errors.New("参数异常")
	}
	return nil, nil
}

func FriendDeleteCheck(ctx *gin.Context, req interface{}) (interface{}, error) {
	reqD := req.(*message.FriendDeleteReq)
	if reqD.ID <= 0 {
		return consts.ParamErrorRespHeader, errors.New("请求参数错误")
	}
	return nil, nil
}
