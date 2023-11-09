package routes

import (
	"JuneGoBlog/internal/api/check"
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/api/server"
	middleware2 "JuneGoBlog/internal/middleware"
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
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneMiddleware.Permiter(middleware2.AdminPermit),
		juneGin.EasyHandler(check.FriendShipUnShowListCheck,
			server.FriendUnShowListLogic, message.FriendUnShowListReq{}),
	}
}

func friendApprovalRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneMiddleware.Permiter(middleware2.AdminPermit),
		juneGin.EasyHandler(check.FriendApprovalCheck,
			server.FriendApprovalLogic, message.FriendApprovalReq{}),
	}
}

func friendDeleteRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
		juneMiddleware.Auth(middleware2.TokenAuth),
		juneMiddleware.Permiter(middleware2.AdminPermit),
		juneGin.EasyHandler(check.FriendDeleteCheck,
			server.FriendDeleteLogic, message.FriendDeleteReq{}),
	}
}

func friendApplicationRoutes() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.NoStoreMiddleware(),
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
