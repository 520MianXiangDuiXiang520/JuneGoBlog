package server

import (
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"fmt"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/gin-gonic/gin"
)

func TagListLogin(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	var resp message.TagListResp
	tags := make([]dao.Tag, 0)
	if err := dao.QueryAllTagsOrderByTime(&tags); err != nil {
		msg := fmt.Sprintf("QueryAllTagsOrderByTime Error!!!")
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Tags = tags
	resp.Total = len(tags)
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func TagAddLogin(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	reqA := req.(*message.TagAddReq)
	var resp message.TagAddResp
	if err := dao.AddTag(reqA.TagName); err != nil {
		msg := fmt.Sprintf("Add Tag Error, name = [%s]\n", reqA.TagName)
		juneLog.ExceptionLog(err, msg)
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}
