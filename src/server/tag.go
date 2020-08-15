package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

func TagListLogin(ctx *gin.Context, req message.BaseReqInter) message.BaseRespInter {
	var resp message.TagListResp
	tags := make([]dao.Tag, 0)
	if err := dao.QueryAllTagsOrderByTime(&tags); err != nil {
		log.Printf("QueryAllTagsOrderByTime Error!!!")
		return consts.SystemErrorRespHeader
	}
	tagInfos := make([]message.TagInfo, 0)
	for _, tag := range tags {
		tagInfos = append(tagInfos, message.TagInfo{
			// TODO: 从 article 表中查
			ArticleTotal: 10,
			Tag:          tag,
		})
	}
	resp.Tags = tagInfos
	resp.Total = len(tagInfos)
	resp.Header = consts.SuccessRespHeader
	return resp
}

func TagAddLogin(ctx *gin.Context, req message.BaseReqInter) message.BaseRespInter {
	reqA := req.(*message.TagAddReq)
	var resp message.TagAddResp
	if err := dao.AddTag(reqA.TagName); err != nil {
		log.Printf("Add Tag Error, name = [%s]\n", reqA.TagName)
		return consts.SystemErrorRespHeader
	}
	resp.Header = consts.SuccessRespHeader
	return resp
}
