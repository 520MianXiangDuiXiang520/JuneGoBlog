package message

type BaseRespHeader struct {
	Code int    `json:"code"`
	Msg string  `json:"msg"`
}
