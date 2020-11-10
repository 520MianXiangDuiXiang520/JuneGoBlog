package util

import (
	"fmt"
	"strings"
)

func GetNotificationTemplate(args map[string]string) string {
	msg := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JuneGoBlog 评论回复通知</title>
</head>
<body>
    <center><h1>JuneGoBlog 评论回复通知</h1></center>

    <center><p>您的文章 <a href="{# articleLink #}">{# articleTitle #}</a> 有了来自 {# talkerName #} 的新评论：</p></center>
    <br>
    <center><p>{# talkText #}</p></center>
</body>
</html>`
	for k, v := range args {
		msg = strings.Replace(msg, fmt.Sprintf("{# %s #}", k), v, -1)
	}
	return msg
}

// args: siteLink, articleLink, sourceTalk, replyTalk
func GetTalkTemplate(args map[string]string) string {
	msg := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JuneGoBlog 评论回复通知</title>
</head>
<body>
    <center><h1>JuneGoBlog 评论回复通知</h1></center>

    <center><p>您在 <a href="{# siteLink #}">JuneGoBlog</a> 中的评论 <a href="{# articleLink #}">{# sourceTalk #}</a> 有了新的回复： {# replyTalk #} 快去看看吧！</p></center>
</body>
</html>`
	for k, v := range args {
		msg = strings.Replace(msg, fmt.Sprintf("{# %s #}", k), v, -1)
	}
	return msg

}
