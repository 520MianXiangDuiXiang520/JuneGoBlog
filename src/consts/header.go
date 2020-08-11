package consts

import (
	"JuneGoBlog/src/message"
	"net/http"
)

// 常用的响应头
var (
	SuccessRespHeader     = message.BaseRespHeader{Code: http.StatusOK, Msg: "ok"}
	SystemErrorRespHeader = message.BaseRespHeader{Code: http.StatusInternalServerError, Msg: "系统异常"}
	ParamErrorRespHeader  = message.BaseRespHeader{Code: http.StatusBadRequest, Msg: "参数错误"}
)

