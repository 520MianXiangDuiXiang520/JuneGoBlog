package middleware

import (
	"JuneGoBlog/internal/db/old"
	"github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	"github.com/gin-gonic/gin"
)

// 使用Token进行认证，如果使用此中间件后请求中 Cookie 未携带 SESSIONID 字段
// 或所携带的值错误，将会返回一个 401 的 HTTP 错误，用于后台接口认证
func TokenAuth(context *gin.Context) (middleware.UserBase, bool) {
	token, err := context.Cookie("SESSIONID")
	if err != nil {
		return nil, false
	}
	user, ok := old.GetUserByToken(token)
	if !ok {
		return nil, false
	}
	return user, true
}

// 用于发表评论时的认证，不会拦截未登录的请求
func TalkingAuth(ctx *gin.Context) (middleware.UserBase, bool) {
	token, err := ctx.Cookie("SESSIONID")
	if err != nil {
		return nil, true
	}
	user, _ := old.GetUserByToken(token)
	return user, true
}
