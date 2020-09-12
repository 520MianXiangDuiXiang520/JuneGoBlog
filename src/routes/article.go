package routes

import (
	"JuneGoBlog/src/check"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", articleListRoutes()...)
	rg.POST("/detail", articleDetailRoutes()...)
	rg.POST("/tags", articleTagsRoutes()...)
}

func articleTagsRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.ArticleTagsCheck,
			server.ArticleTagsLogic, &message.ArticleTagsReq{}),
	}
}

func articleDetailRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.ArticleDetailCheck,
			server.ArticleDetailLogic, &message.ArticleDetailReq{}),
	}
}

func articleListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		junebao_top.EasyHandler(check.ArticleListCheck,
			server.ArticleListLogic, &message.ArticleListReq{}),
	}
}
