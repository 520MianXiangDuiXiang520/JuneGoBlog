package routes

import (
	"JuneGoBlog/internal/api/check"
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/api/server"
	"JuneGoBlog/internal/middleware"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneMiddleware "github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRegister(rg *gin.RouterGroup) {
	rg.POST("/list", articleListRoutes()...)
	rg.POST("/detail", articleDetailRoutes()...)
	rg.POST("/tags", articleTagsRoutes()...)
	rg.POST("/add", articleAddRoutes()...)
	rg.POST("/update", articleUpdateRoutes()...)
	rg.POST("/delete", articleDeleteRoutes()...)
}

func articleTagsRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneGin.EasyHandler(check.ArticleTagsCheck,
			server.ArticleTagsLogic, message.ArticleTagsReq{}),
	}
}

func articleDetailRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.ArticleDetailCheck,
			server.ArticleDetailLogic, message.ArticleDetailReq{}),
	}
}

func articleListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneGin.EasyHandler(check.ArticleListCheck,
			server.ArticleListLogic, message.ArticleListReq{}),
	}
}

func articleAddRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.ArticleAddCheck,
			server.ArticleAddLogic, message.ArticleAddReq{}),
	}
}

func articleUpdateRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.ArticleUpdateCheck,
			server.ArticleUpdateLogic, message.ArticleUpdateReq{}),
	}
}

func articleDeleteRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.ArticleDeleteCheck,
			server.ArticleDeleteLogic, message.ArticleDeleteReq{}),
	}
}
