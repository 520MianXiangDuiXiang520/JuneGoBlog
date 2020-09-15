package src

import (
	"github.com/gin-gonic/gin"
)

// 管理员权限
func AdminPermit(ctx *gin.Context) bool {
	// sessionID, err := ctx.Cookie("SESSIONID")
	// if err != nil {
	//     return false
	// }
	// user, ok := dao.GetUserByToken(sessionID)
	// if !ok {
	//     return false
	// }
	return true
}
