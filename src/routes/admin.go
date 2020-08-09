package routes

import (
	"github.com/gin-gonic/gin"
)

func AdminRegister(rg *gin.RouterGroup) {
	rg.POST("/", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"msg": "admin",
		})
	})
}
