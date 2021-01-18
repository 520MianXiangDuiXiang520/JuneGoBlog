package routes

import (
	"JuneGoBlog/src/check"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/middleware"
	"JuneGoBlog/src/server"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneMiddleware "github.com/520MianXiangDuiXiang520/GinTools/gin/middleware"
	"github.com/gin-gonic/gin"
)

func FriendShipRoutes(rg *gin.RouterGroup) {
	rg.POST("/list", friendListRoutes()...)
	rg.POST("/unshow", friendshipUnShowRoutes()...)       // 申请中的友链
	rg.POST("/application", friendApplicationRoutes()...) // 申请
	rg.POST("/delete", friendDeleteRoutes()...)           // 需要管理员权限
	rg.POST("/approval", friendApprovalRoutes()...)       // 审批，需要管理员权限
}

func friendshipUnShowRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.FriendShipUnShowListCheck,
			server.FriendUnShowListLogic, message.FriendUnShowListReq{}),
	}
}

func friendApprovalRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.FriendApprovalCheck,
			server.FriendApprovalLogic, message.FriendApprovalReq{}),
	}
}

func friendDeleteRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware.TokenAuth),
		juneMiddleware.Permiter(middleware.AdminPermit),
		juneGin.EasyHandler(check.FriendDeleteCheck,
			server.FriendDeleteLogic, message.FriendDeleteReq{}),
	}
}

func friendApplicationRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NoStoreMiddleware(),
		juneGin.EasyHandler(check.FriendApplicationCheck,
			server.FriendApplicationLogic, message.FriendApplicationReq{}),
	}
}

func friendListRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		juneGin.EasyHandler(check.FriendShipListCheck,
			server.FriendShipListLogic, message.FriendShipListReq{}),
	}
}
