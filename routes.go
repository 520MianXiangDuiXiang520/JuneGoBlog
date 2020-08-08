package main

import (
	"JuneGoBlog/src/admin"
	"JuneGoBlog/src/article"
	"JuneGoBlog/src/tag"
	"JuneGoBlog/src/talking"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)


func Register(c *gin.Engine) {
	utils.HandlerRoute(c, "api/article", article.Register)
	utils.HandlerRoute(c, "api/tag", tag.Register)
	utils.HandlerRoute(c, "api/talking", talking.Register)
	utils.HandlerRoute(c, "api/admin", admin.Register)
}
