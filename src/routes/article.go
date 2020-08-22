package routes

import (
	"JuneGoBlog/src/check"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", ArticleListRoutes()...)
	rg.POST("/detail", articleDetailRoutes()...)
}

func articleDetailRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.ArticleDetailCheck,
			server.ArticleDetailLogic, &message.ArticleDetailReq{}),
	}
}

func ArticleListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.ArticleListCheck,
			server.ArticleListLogic, &message.ArticleListReq{}),
	}
}
