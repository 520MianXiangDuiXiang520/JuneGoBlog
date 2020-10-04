package server

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/junebao.top/utils"
	"JuneGoBlog/src/message"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"time"
)

func getArticleTagsInfo(id int) ([]message.TagInfo, error) {
	tags := make([]dao.Tag, 0)
	tagsInfoList := make([]message.TagInfo, 0)
	if err := dao.QueryAllTagsByArticleID(id, &tags); err != nil {
		msg := fmt.Sprintf("get all tags by article id fail, %v", id)
		utils.ExceptionLog(err, msg)
		return tagsInfoList, err
	}

	for _, tagInfo := range tags {
		articleTotal, err := dao.QueryArticleTotalByTagID(tagInfo.ID)
		if err != nil {
			mes := fmt.Sprintf("query article total by cache fail !")
			utils.ExceptionLog(err, mes)
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
		utils.ExceptionLog(err, mes)
		return junebaotop.SystemErrorRespHeader
	}
	resp.ID = reqL.ArticleID
	resp.Tags = tags
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

func articleListByTag(tagID, page, pageSize int) (*message.ArticleListResp, error) {

	articleList, total, err := dao.QueryArticleInfoByLimitWithTag(tagID, page, pageSize)
	if err != nil {
		return nil, err
	}
	return &message.ArticleListResp{
		Header:      junebaotop.SuccessRespHeader,
		ArticleList: articleList,
		Total:       total,
	}, nil
}

// 文章列表逻辑
func ArticleListLogic(ctx *gin.Context,
	req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	reqL := req.(*message.ArticleListReq)
	resp := message.ArticleListResp{}

	if reqL.Tag != 0 {
		response, err := articleListByTag(reqL.Tag, reqL.Page, reqL.PageSize)
		if err != nil {
			return junebaotop.SystemErrorRespHeader
		}
		return response
	}

	articleList, total, err := dao.QueryArticleInfoByLimit(reqL.Page, reqL.PageSize)
	if err != nil {
		return junebaotop.SystemErrorRespHeader
	}
	resp.ArticleList = articleList
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

func ArticleAddLogic(ctx *gin.Context,
	req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.ArticleAddReq)
	resp := message.ArticleAddResp{}
	user, ok := ctx.Get("user")
	if !ok {
		return junebaotop.UnauthorizedRespHeader
	}
	author := user.(*dao.User)

	if len(request.Abstract) <= 0 {
		request.Abstract = getAbstract(request.Text)
	}

	newArticle := dao.Article{
		Text:       request.Text,
		Title:      request.Title,
		AuthorID:   author.ID,
		Abstract:   request.Abstract,
		CreateTime: time.Now(),
	}
	_, err := dao.AddArticle(&newArticle, request.Tags)
	if err != nil {
		return junebaotop.SystemErrorRespHeader
	}
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

func getAbstract(text string) string {
	abstractList := strings.Split(text, consts.AbstractSplitStr)
	sp := src.Setting.AbstractLen
	if sp > len(text) {
		sp = len(text)
	}
	if len(abstractList) < 2 {
		return text[:sp]
	}
	// 避免标题加入摘要
	r := regexp.MustCompile("^#(.|\\r|)+").ReplaceAllString(abstractList[0], "")
	r = strings.Replace(r, "\n", "", len(r))
	return r
}

func ArticleUpdateLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.ArticleUpdateReq)
	resp := message.ArticleUpdateResp{}

	user, ok := ctx.Get("user")
	if !ok {
		return junebaotop.UnauthorizedRespHeader
	}
	author := user.(*dao.User)

	if request.CreateTime.Unix() < 0 {
		request.CreateTime = time.Now()
	}

	abstract := request.Abstract
	if abstract == "" {
		abstract = getAbstract(request.Text)
	}

	// update article table
	err := dao.UpdateArticle(request.ID, &dao.Article{
		ID:         request.ID,
		Text:       request.Text,
		Title:      request.Title,
		AuthorID:   author.ID,
		Abstract:   abstract,
		CreateTime: request.CreateTime,
	})
	if err != nil {
		return junebaotop.SystemErrorRespHeader
	}
	// update article_tag table
	err = dao.UpdateArticleTagsByIntList(request.ID, request.Tags)
	if err != nil {
		return junebaotop.SystemErrorRespHeader
	}

	resp.Header = junebaotop.SuccessRespHeader
	return resp
}

func ArticleDeleteLogic(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
	request := req.(*message.ArticleDeleteReq)
	resp := message.ArticleDeleteResp{}
	err := dao.DeleteArticle(request.ID)
	if err != nil {
		return junebaotop.SystemErrorRespHeader
	}
	resp.Header = junebaotop.SuccessRespHeader
	return resp
}
