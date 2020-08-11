package server

import (
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
)

func ListLogic(ctx *gin.Context, req interface{}) interface{} {

	var resp *message.ArticleListResp

	_ = ctx.ShouldBindJSON(&req)
	// TODO: dao

	articleList := make([]dao.ArticleInfo, 0)
	articleInfo := dao.ArticleInfo{Name: "Go 内存模型",
		CreateTime: 12345678, Abstract: "xxxx", Text: "Hello Go", ReadingAmount: 10,
		Tags: []dao.TagInfo{}}
	articleList = append(articleList, articleInfo)
	resp = &message.ArticleListResp{Header: consts.SuccessRespHeader, ArticleList: articleList, Total: 1}
	return resp
}
