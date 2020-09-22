package middleware

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
)

// 管理员权限
func AdminPermit(ctx *gin.Context) bool {
	user, ok := ctx.Get("user")
	if !ok {
		return false
	}
	u := user.(*dao.User)
	return u.Permiter == consts.AdminPermission
}
