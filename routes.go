package main

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/routes"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneMiddle "github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Engine) {
	c.Use(middleware.ApiView(), juneMiddle.CorsHandler(src.GetSetting().CorsAccessList))
	juneGin.URLPatterns(c, "api/article", routes.ArticleRegister)
	juneGin.URLPatterns(c, "api/tag", routes.TagRegister)
	juneGin.URLPatterns(c, "api/talking", routes.TalkingRegister)
	juneGin.URLPatterns(c, "api/admin", routes.AdminRegister)
	juneGin.URLPatterns(c, "api/friendship", routes.FriendShipRoutes)
	juneGin.URLPatterns(c, "api/auth", routes.AuthRegister)

}
