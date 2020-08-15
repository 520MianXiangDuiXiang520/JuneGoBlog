package middleware

import "github.com/gin-gonic/gin"

// 拦截 GET 请求
func ApiView() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == "GET" {
			context.Abort()
			context.Writer.WriteHeader(200)
			_, _ = context.Writer.WriteString("           *   Welcome JuneBlog   *\n\n     _                  ____  _             \n    | |_   _ _ __   ___| __ )| | ___   __ _ \n _  | | | | | '_ \\ / _ \\  _ \\| |/ _ \\ / _` |\n| |_| | |_| | | | |  __/ |_) | | (_) | (_| |\n \\___/ \\__,_|_| |_|\\___|____/|_|\\___/ \\__, |\n                                      |___/")
		}
	}
}
