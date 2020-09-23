package middleware

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	junebao_top "JuneGoBlog/src/junebao.top"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 管理员权限，允许访客查看，除非未登录，否则不会响应 403
func AdminPermit(ctx *gin.Context) bool {
	user, ok := ctx.Get("user")
	if !ok {
		return false
	}
	u := user.(*dao.User)
	// 如果用户持有游客身份，直接响应 200 不再去执行业务逻辑
	if u.Permiter != consts.AdminPermission {
		ctx.Abort()
		ctx.JSON(http.StatusOK,
			junebao_top.SuccessRespHeader)
	}
	return true
}
