package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/utils"
	"github.com/gin-gonic/gin"
)

func ListLogic(ctx *gin.Context) utils.RespHeader {
	var req message.ListReq
	var resp *message.ListResp

	_ = ctx.ShouldBindJSON(&req)
	// TODO: dao

	articleList := make([]dao.ArticleInfo, 0)
	articleInfo := dao.ArticleInfo{Name: "Go 内存模型",
		CreateTime: 12345678, Abstract: "xxxx", Text: "Hello Go", ReadingAmount: 10,
		Tags: []dao.TagInfo{}}
	articleList = append(articleList, articleInfo)
	resp = &message.ListResp{Header: consts.SuccessHead, ArticleList: articleList, Total: 1}
	return resp
}
