package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func TalkingListLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.TalkingListReq)
	resp := message.TalkingListResp{}
	talks, err := dao.QueryTalksByArticleIDLimit(request.ArticleID, request.Page, request.PageSize)
	if err != nil {
		msg := fmt.Sprintf("Fail to query talks, request is %v ", request)
		util.ExceptionLog(err, msg)
	}
	resp.HasNext = true
	if len(talks) < request.PageSize {
		resp.HasNext = false
	}
	resp.Talks = talks
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

func TalkingAddLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.TalkingAddReq)
	resp := message.TalkingAddResp{}
	if !dao.HasArticle(request.ArticleID) {
		return junebaotop.ParamErrorRespHeader
	}
	if request.Type == consts.ChildTalkType {
		if !dao.HasTalk(request.PTalkID) {
			return junebaotop.ParamErrorRespHeader
		}
	} else {
		request.PTalkID = 0
	}
	if request.Username == "" {
		request.Username = strings.Split(request.Email, "@")[0]
	}
	err := dao.AddTalk(&dao.Talks{
		ArticleID:  request.ArticleID,
		Text:       request.Text,
		Username:   request.Username,
		PTalkID:    request.PTalkID,
		Email:      request.Email,
		Type:       request.Type,
		SiteLink:   request.SiteLink,
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		msg := fmt.Sprintf("Fail to add new talk, request = %v", request)
		util.ExceptionLog(err, msg)
		return junebaotop.SystemErrorRespHeader
	}
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}
