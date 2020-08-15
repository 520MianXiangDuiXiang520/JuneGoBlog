package src

import "github.com/gin-gonic/gin"

// 管理员权限
func AdminPermit(ctx *gin.Context) bool {
	return true
}
