package middleware

import (
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/junebao.top/middleware"
	"github.com/gin-gonic/gin"
)

func TokenAuth(context *gin.Context) (middleware.UserBase, bool) {
	token, err := context.Cookie("SESSIONID")
	if err != nil {
		return nil, false
	}
	user, ok := dao.GetUserByToken(token)
	if !ok {
		return nil, false
	}
	return user, true
}
