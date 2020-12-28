package check

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"errors"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TalkingListCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.TalkingListReq)
	if !dao.HasArticle(request.ArticleID) {
		return juneGin.ParamErrorRespHeader, errors.New("can not find this article")
	}
	if request.ParentTalkID > 0 && !dao.HasTalk(request.ParentTalkID) {
		return juneGin.ParamErrorRespHeader, errors.New("can not find this talk")
	}
	return http.StatusOK, nil
}

func TalkingAddCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.TalkingAddReq)
	if request.Type != consts.RootTalkType && request.Type != consts.ChildTalkType {
		return juneGin.ParamErrorRespHeader, errors.New("wrong type")
	}
	return http.StatusOK, nil
}
