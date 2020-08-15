package junebao_top

import (
	"github.com/gin-gonic/gin"
)

type DoChildRouteFunc func(g *gin.RouterGroup) // 子路由

type Route interface {
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

// URL 调度器
func URLPatterns(route Route, path string, childRoute DoChildRouteFunc, middles ...gin.HandlerFunc) {
	group := route.Group(path)
	for _, mid := range middles {
		group.Use(mid)
	}
	childRoute(group)
}
