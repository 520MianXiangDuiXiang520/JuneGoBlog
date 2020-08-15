package middleware

import (
	"JuneGoBlog/src/consts"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 将一个 License 包装成中间件， 如果 license.Authorization
// 为 false， 响应 403
func Permiter(license License) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !license.Authorization(context) {
			context.Abort()
			log.Printf("Abort")
			context.JSON(http.StatusForbidden, consts.ForbiddenErrorRespHeader)
		}
	}
}

type License interface {
	Authorization(ctx *gin.Context) bool
}

// 管理员权限
type AdminPermit struct{}

func (pl AdminPermit) Authorization(ctx *gin.Context) bool {
	return true
}
