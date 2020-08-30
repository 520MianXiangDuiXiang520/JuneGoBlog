package main

import (
	"JuneGoBlog/src/junebao.top"
	middleware2 "JuneGoBlog/src/junebao.top/middleware"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/routes"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Engine) {
	c.Use(middleware.ApiView(), middleware2.CorsHandler())
	junebao_top.URLPatterns(c, "api/article", routes.ArticleRegister)
	junebao_top.URLPatterns(c, "api/tag", routes.TagRegister)
	junebao_top.URLPatterns(c, "api/talking", routes.TalkingRegister)
	junebao_top.URLPatterns(c, "api/admin", routes.AdminRegister)
	junebao_top.URLPatterns(c, "api/friendship", routes.FriendShipRoutes)

}
