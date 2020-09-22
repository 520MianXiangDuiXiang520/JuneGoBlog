package routes

import (
	"JuneGoBlog/src/check"
	junebao_top "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/junebao.top/middleware"
	"JuneGoBlog/src/message"
	middleware2 "JuneGoBlog/src/middleware"
	"JuneGoBlog/src/server"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", articleListRoutes()...)
	rg.POST("/detail", articleDetailRoutes()...)
	rg.POST("/tags", articleTagsRoutes()...)
	rg.POST("/add", articleAddRoutes()...)
	rg.POST("/update", articleUpdateRoutes()...)
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

func articleAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		junebao_top.EasyHandler(check.ArticleAddCheck,
			server.ArticleAddLogic, &message.ArticleAddReq{}),
	}
}

func articleUpdateRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		junebao_top.EasyHandler(check.ArticleUpdateCheck,
			server.ArticleUpdateLogic, &message.ArticleUpdateReq{}),
	}
}
