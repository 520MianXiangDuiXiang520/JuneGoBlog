package routes

import (
	"github.com/gin-gonic/gin"
)

func TagRegister(rg *gin.RouterGroup) {
	rg.POST("/list", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
}
