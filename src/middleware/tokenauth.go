package middleware

import (
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
)

func TokenAuth(context *gin.Context) (interface{}, bool) {
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
