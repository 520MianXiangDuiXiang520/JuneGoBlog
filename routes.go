package main

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/routes"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneMiddle "github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	middleware2 "github.com/520MianXiangDuiXiang520/ginUtils/middleware"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Engine) {
	c.Use(
		juneMiddle.CorsHandler(src.GetSetting().CorsAccessList),
		middleware2.Throttled(middleware2.SimpleThrottle(
			middleware2.ThrottledRuleByUserAgentAndIP, "30/m")),
		middleware.ApiView(),
	)
	juneGin.URLPatterns(c, "api/article", routes.ArticleRegister)
	juneGin.URLPatterns(c, "api/tag", routes.TagRegister)
	juneGin.URLPatterns(c, "api/talking", routes.TalkingRegister)
	juneGin.URLPatterns(c, "api/friendship", routes.FriendShipRoutes)
	juneGin.URLPatterns(c, "api/auth", routes.AuthRegister)

}
