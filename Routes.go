package JuneGoBlog

import "github.com/gin-gonic/gin"

func Routes(e *gin.Engine) {
	e.GET("/", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg": "ok",
		})
	})
}
