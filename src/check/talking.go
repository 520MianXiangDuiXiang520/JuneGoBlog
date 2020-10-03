package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TalkingListCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	request := req.(*message.TalkingListReq)
	if !dao.HasArticle(request.ArticleID) {
		return junebao_top.ParamErrorRespHeader, errors.New("can not find this article")
	}
	if request.PageSize <= 0 || request.Page <= 0 {
		return junebao_top.ParamErrorRespHeader, errors.New("page or page-size is wrong")
	}
	if request.ParentTalkID > 0 && !dao.HasTalk(request.ParentTalkID) {
		return junebao_top.ParamErrorRespHeader, errors.New("can not find this talk")
	}
	return http.StatusOK, nil
}

func TalkingAddCheck(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
	request := req.(*message.TalkingAddReq)
	if len(request.Text) <= 0 {
		return junebao_top.ParamErrorRespHeader, errors.New("text is none")
	}
	if !util.IsEmail(request.Email) {
		return junebao_top.ParamErrorRespHeader, errors.New("email has wrong format")
	}
	if request.Type != consts.RootTalkType && request.Type != consts.ChildTalkType {
		return junebao_top.ParamErrorRespHeader, errors.New("wrong type")
	}
	return http.StatusOK, nil
}
