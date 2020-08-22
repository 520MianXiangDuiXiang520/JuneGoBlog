package server

import (
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

func TagListLogin(ctx *gin.Context, req junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	var resp message.TagListResp
	tags := make([]dao.Tag, 0)
	if err := dao.QueryAllTagsOrderByTime(&tags); err != nil {
		log.Printf("QueryAllTagsOrderByTime Error!!!")
		return junebao_top.SystemErrorRespHeader
	}
	tagInfos := make([]message.TagInfo, 0)
	for _, tag := range tags {
		tagInfos = append(tagInfos, message.TagInfo{
			// TODO: 从 article 表中查
			ArticleTotal: 10,
			Name:         tag.Name,
			CreateTime:   tag.CreateTime.Unix(),
		})
	}
	resp.Tags = tagInfos
	resp.Total = len(tagInfos)
	resp.Header = junebao_top.SuccessRespHeader
	return resp
}

func TagAddLogin(ctx *gin.Context, req junebao_top.BaseReqInter) junebao_top.BaseRespInter {
	reqA := req.(*message.TagAddReq)
	var resp message.TagAddResp
	if err := dao.AddTag(reqA.TagName); err != nil {
		log.Printf("Add Tag Error, name = [%s]\n", reqA.TagName)
		return junebao_top.SystemErrorRespHeader
	}
	resp.Header = junebao_top.SuccessRespHeader
	return resp
}
