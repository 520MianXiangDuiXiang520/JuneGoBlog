package consts

import (
	"JuneGoBlog/src/message"
	"net/http"
)

var  SuccessHead = message.BaseRespHeader{Code: http.StatusOK, Msg: "ok"}
var SystemError = message.BaseRespHeader{Code: http.StatusInternalServerError, Msg: "系统异常"}
