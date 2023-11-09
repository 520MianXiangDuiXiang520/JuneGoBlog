package middleware

import (
	"JuneGoBlog/internal/consts"
	"JuneGoBlog/internal/db/old"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 管理员权限，允许访客查看，除非未登录，否则不会响应 403
func AdminPermit(ctx *gin.Context) bool {
	user, ok := ctx.Get("user")
	if !ok {
		return false
	}
	u := user.(*old.User)
	// 如果用户持有游客身份，直接响应 200 不再去执行业务逻辑
	if u.Permiter != consts.AdminPermission {
		ctx.Abort()
		ctx.JSON(http.StatusOK,
			juneGin.SuccessRespHeader)
	}
	return true
}
