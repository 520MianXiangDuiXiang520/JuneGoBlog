package routes

import (
	"JuneGoBlog/src/dao"
	middleware2 "JuneGoBlog/src/junebao.top/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRegister(rg *gin.RouterGroup) {
	rg.Use(middleware2.Auth(func(context *gin.Context) (interface{}, bool) {
		return dao.User{
			ID:       1,
			Username: "zhangsan",
			Password: "1111",
			Permiter: 1,
		}, true
	}))
	rg.POST("/", func(context *gin.Context) {
		user, _ := context.Get("user")
		context.JSON(200, map[string]interface{}{
			"msg": user,
		})
	})
}
