package server

import (
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func getArticleTagsInfo(id int) ([]message.TagInfo, error) {
	tags := make([]dao.Tag, 0)
	tagsInfoList := make([]message.TagInfo, 0)
	if err := dao.QueryAllTagsByArticleID(id, &tags); err != nil {
		msg := fmt.Sprintf("get all tags by article id fail, %v", id)
		util.ExceptionLog(err, msg)
		return tagsInfoList, err
	}

	for _, tagInfo := range tags {
		articleTotal, err := dao.QueryArticleTotalByTagID(tagInfo.ID)
		if err != nil {
			mes := fmt.Sprintf("query article total by cache fail !")
			util.ExceptionLog(err, mes)
			return nil, err
		}
		tagsInfoList = append(tagsInfoList, message.TagInfo{
			ID:           tagInfo.ID,
			Name:         tagInfo.Name,
			CreateTime:   tagInfo.CreateTime.Unix(),
			ArticleTotal: articleTotal,
		})
	}
	return tagsInfoList, nil
}

func ArticleTagsLogic(ctx *gin.Context,
	req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	reqL := req.(*message.ArticleTagsReq)
	resp := message.ArticleTagsResp{}
	tags, err := getArticleTagsInfo(reqL.ArticleID)
	if err != nil {
		mes := fmt.Sprintf("get article tags fail, "+
			"article id = %v ", reqL.ArticleID)
		util.ExceptionLog(err, mes)
		return junebaotop.SystemErrorRespHeader
	}
	resp.ID = reqL.ArticleID
	resp.Tags = tags
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

// 文章列表逻辑
func ArticleListLogic(ctx *gin.Context,
	req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	reqL := req.(*message.ArticleListReq)
	resp := message.ArticleListResp{}
	total, err := dao.QueryArticleTotal()
	if err != nil {
		log.Printf("获取文章总数失败！")
		return junebaotop.SystemErrorRespHeader
	}

	articleList, err := dao.QueryArticleInfoByLimit(reqL.Page, reqL.PageSize, total)
	if err != nil {
		log.Printf("QueryArticleInfoByLimit Error !")
		return junebaotop.SystemErrorRespHeader
	}
	resp.ArticleList = make([]dao.ArticleInfo, 0)
	for _, article := range articleList {
		tags := make([]dao.Tag, 0)
		err := dao.QueryAllTagsByArticleID(article.ID, &tags)
		if err != nil {
			msg := fmt.Sprintf("get article tags fail, article id = %v", article.ID)
			util.ExceptionLog(err, msg)
		}
		resp.ArticleList = append(resp.ArticleList, dao.ArticleInfo{
			Tags:       tags,
			ID:         article.ID,
			Title:      article.Title,
			CreateTime: article.CreateTime,
			Abstract:   article.Abstract,
			Author:     "Junebao",
		})
	}
	resp.Header = junebaotop.SuccessRespHeader
	resp.Total = total
	return resp
}

func ArticleDetailLogic(ctx *gin.Context,
	req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	resp := message.ArticleDetailResp{}
	reqD := req.(*message.ArticleDetailReq)

	article, _ := dao.QueryArticleDetail(reqD.ArticleID)
	resp.ID = article.ID
	resp.Text = article.Text
	resp.CreateTime = article.CreateTime
	resp.Abstract = article.Abstract
	resp.AuthorID = article.AuthorID
	resp.Title = article.Title
	resp.BaseRespHeader = junebaotop.SuccessRespHeader
	return resp
}

func ArticleAddLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.ArticleAddReq)
	resp := message.ArticleAddResp{}
	// TODO:...
	log.Println(request)
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}
