package check

import (
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FriendShipListCheck(ctx *gin.Context) (utils.RespHeader, error) {
	// 无请求参数，不需要校验
	return http.StatusOK, nil
}
