package middleware

import (
	junebaotop "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "http://localhost:8081") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,"+
			"session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT,"+
			" X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type,"+
			" Pragma,token,openid,opentoken, cookies, Cookies, cookie, Cookies")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Set("content-type", "application/json")
		//设置返回格式是json

		if method == "OPTIONS" {
			context.Abort()
			context.JSON(http.StatusOK, junebaotop.SuccessRespHeader)
		}
		context.Next()
	}
}
