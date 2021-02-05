package server

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"fmt"
	juneEmail "github.com/520MianXiangDuiXiang520/GinTools/email"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func TalkingListLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.TalkingListReq)
	resp := message.TalkingListResp{}
	talks, err := dao.QueryTalksByArticleIDLimit(request.ArticleID, request.Page, request.PageSize)
	if err != nil {
		msg := fmt.Sprintf("Fail to query talks, request is %v ", request)
		juneLog.ExceptionLog(err, msg)
	}
	resp.HasNext = true
	if len(talks) < request.PageSize {
		resp.HasNext = false
	}
	resp.Talks = talks
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func TalkingAddLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.TalkingAddReq)
	resp := message.TalkingAddResp{}
	if !dao.HasArticle(request.ArticleID) {
		return juneGin.ParamErrorRespHeader
	}
	if request.Type == consts.ChildTalkType {
		if !dao.HasTalk(request.PTalkID) {
			return juneGin.ParamErrorRespHeader
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
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	// 发送邮件通知
	go func(r *message.TalkingAddReq) {
		if r.PTalkID != 0 {
			err = sendNotification(r)
			if err != nil {
				msg := fmt.Sprintf("send email error! %v", r)
				juneLog.ExceptionLog(err, msg)
			}
		}
		err = sendNotificationToAuthor(r)
		if err != nil {
			msg := fmt.Sprintf("send email error! %v", r)
			juneLog.ExceptionLog(err, msg)
		}
	}(request)
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func sendNotificationToAuthor(r *message.TalkingAddReq) error {
	subject := "文章有了新评论"
	article, _ := dao.QueryArticleByID(r.ArticleID)
	body := util.GetNotificationTemplate(map[string]string{
		"articleTitle": article.Title,
		"talkerName":   r.Username,
		"talkText":     r.Text,
		"articleLink":  fmt.Sprintf("%s%d", src.GetSetting().Others.DetailLink, r.ArticleID),
	})
	return juneEmail.Send(&juneEmail.Context{
		ToList: []juneEmail.Role{
			{Address: src.GetSetting().Others.MyEmail},
		},
		Subject: subject,
		Body:    body,
	})
}

func sendNotification(r *message.TalkingAddReq) error {
	subject := "JuneBlog 评论回复通知"
	st, _ := dao.QueryTalkByTalkID(r.PTalkID)
	body := util.GetTalkTemplate(map[string]string{
		"siteLink":    src.GetSetting().Others.SiteLink,
		"articleLink": fmt.Sprintf("%s%d", src.GetSetting().Others.DetailLink, r.ArticleID),
		"sourceTalk":  st.Text,
		"replyTalk":   r.Text,
	})
	return juneEmail.Send(&juneEmail.Context{
		ToList: []juneEmail.Role{
			{Address: st.Email, Name: st.Username},
		},
		Subject: subject,
		Body:    body,
	})
}
