package server

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	email "JuneGoBlog/src/junebao.top/email"
	"JuneGoBlog/src/junebao.top/utils"
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
		utils.ExceptionLog(err, msg)
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
		utils.ExceptionLog(err, msg)
		return junebaotop.SystemErrorRespHeader
	}
	// 发送邮件通知
	go func(r *message.TalkingAddReq) {
		if r.PTalkID != 0 {
			err = sendNotification(r)
			if err != nil {
				msg := fmt.Sprintf("send email error! %v", r)
				utils.ExceptionLog(err, msg)
			}
		}
		err = sendNotificationToAuthor(r)
		if err != nil {
			msg := fmt.Sprintf("send email error! %v", r)
			utils.ExceptionLog(err, msg)
		}
	}(request)
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

func sendNotificationToAuthor(r *message.TalkingAddReq) error {
	subject := "文章有了新评论"
	article, _ := dao.QueryArticleByID(r.ArticleID)
	body := util.GetNotificationTemplate(map[string]string{
		"articleTitle": article.Title,
		"talkerName":   r.Username,
		"talkText":     r.Text,
		"articleLink":  fmt.Sprintf("http://39.106.168.39/#/detail/%d", r.ArticleID),
	})
	return email.Send(subject, body, []string{src.Setting.MyEmail})
}

func sendNotification(r *message.TalkingAddReq) error {
	subject := "JuneGoBlog 评论回复通知"
	st, _ := dao.QueryTalkByTalkID(r.PTalkID)
	body := util.GetTalkTemplate(map[string]string{
		"siteLink":    "https://junebao.top",
		"articleLink": fmt.Sprintf("http://39.106.168.39/#/detail/%d", r.ArticleID),
		"sourceTalk":  st.Text,
		"replyTalk":   r.Text,
	})
	return email.Send(subject, body, []string{st.Email})
}
