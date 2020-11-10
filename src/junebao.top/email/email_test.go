package email

import (
	"JuneGoBlog/src/util"
	"testing"
)

func TestSimpleEmail_Send(t *testing.T) {

	tos := []string{
		"1771795643@qq.com", "3176869767@qq.com",
	}
	subject := "测试邮件"
	body := "<h1> 测试 <h1>"
	Send(subject, body, tos)
}

func TestSimpleEmail_Send2(t *testing.T) {

	tos := []string{
		"1771795643@qq.com",
	}
	subject := "测试邮件"
	body := util.GetTalkTemplate(map[string]string{
		"siteLink":    "https://junebao.top",
		"articleLink": "http://39.106.168.39/#/detail/67",
		"sourceTalk":  "aaa, aw",
		"replyTalk":   "刷的一下，很快啊！",
	})
	Send(subject, body, tos)
}
