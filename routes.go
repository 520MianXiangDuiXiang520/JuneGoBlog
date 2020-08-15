package main

import (
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/routes"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Engine) {
	c.Use(middleware.ApiView())
	utils.HandlerRoute(c, "api/article", routes.ArticleRegister)
	utils.HandlerRoute(c, "api/tag", routes.TagRegister)
	utils.HandlerRoute(c, "api/talking", routes.TalkingRegister)
	utils.HandlerRoute(c, "api/admin", routes.AdminRegister)
	utils.HandlerRoute(c, "api/friendship", routes.FriendShipRoutes)

}
