package admin

import (
	"JuneGoBlog/src/article"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup) {
	utils.HandlerRoute(rg, "/article", article.Register)
	rg.POST("/", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"msg": "admin",
		})
	})
}
