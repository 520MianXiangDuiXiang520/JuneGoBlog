package consts

import (
	"JuneGoBlog/src/message"
	"net/http"
)

var  SuccessHead = message.BaseRespHeader{Code: http.StatusOK, Msg: "ok"}
