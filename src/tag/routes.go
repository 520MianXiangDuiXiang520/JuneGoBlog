package tag

import (
	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup) {
	rg.POST("/list", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
}
