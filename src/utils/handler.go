package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ChildRoute func(g *gin.RouterGroup)   // 子路由
type CheckReqArgs func(c *gin.Context) int // 检查请求参数
type DoLogic func(c *gin.Context)          // 处理具体的也业务逻辑

type Route interface {
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

// 转发到子路由
func HandlerRoute(route Route, path string, childRoute ChildRoute, middles ...gin.HandlerFunc) {
	fmt.Println("HandlerRoute")
	group := route.Group(path)
	for _, mid := range middles {
		group.Use(mid)
	}
	childRoute(group)
}

