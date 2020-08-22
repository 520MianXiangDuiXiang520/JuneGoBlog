package server

import (
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"log"
)

func getArticleTagsInfo(id int) []message.TagInfo {
	tags := make([]dao.Tag, 0)
	tagsInfoList := make([]message.TagInfo, 0)
	if err := dao.QueryAllTagsByArticleID(id, &tags); err != nil {
		log.Printf("QueryAllTagsByArticleID Error !")
		return tagsInfoList
	}

	for _, tagInfo := range tags {
		articleTotal := dao.QueryArticleTotalByTagID(tagInfo.ID)
		tagsInfoList = append(tagsInfoList, message.TagInfo{
			ID:           tagInfo.ID,
			Name:         tagInfo.Name,
			CreateTime:   tagInfo.CreateTime.Unix(),
			ArticleTotal: articleTotal,
		})
	}
	return tagsInfoList
}

func ArticleListLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	reqL := req.(*message.ArticleListReq)
	resp := message.ArticleListResp{}
	articleList, err := dao.QueryArticleByLimit(reqL.Page, reqL.PageSize)
	if err != nil {
		log.Printf("QueryArticleByLimit Error !")
		return junebaotop.SystemErrorRespHeader
	}
	articleTagsList := make([]message.ArticleTagInfo, 0)
	for _, article := range articleList {
		tagsInfoList := getArticleTagsInfo(article.ID)
		articleTagsList = append(articleTagsList, message.ArticleTagInfo{
			ID:         article.ID,
			AuthorID:   article.AuthorID,
			Title:      article.Title,
			CreateTime: article.CreateTime.Unix(),
			Abstract:   article.Abstract,
			Tags:       tagsInfoList,
		})
	}

	resp.Header = junebaotop.SuccessRespHeader
	resp.ArticleList = articleTagsList
	resp.Total = len(articleList)
	return resp
}

func ArticleDetailLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	resp := message.ArticleDetailResp{}
	reqD := req.(*message.ArticleDetailReq)

	article, _ := dao.QueryArticleDetail(reqD.ArticleID)
	tagsInfoList := getArticleTagsInfo(article.ID)
	resp.ID = article.ID
	resp.Text = article.Text
	resp.CreateTime = article.CreateTime.Unix()
	resp.Abstract = article.Abstract
	resp.AuthorID = article.AuthorID
	resp.Title = article.Title
	resp.Tags = tagsInfoList
	resp.BaseRespHeader = junebaotop.SuccessRespHeader
	return resp
}
