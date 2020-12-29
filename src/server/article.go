package server

import (
	"JuneGoBlog/src"
	"JuneGoBlog/src/consts"
	"JuneGoBlog/src/dao"
	"JuneGoBlog/src/message"
	"JuneGoBlog/src/util"
	"fmt"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	juneLog "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"unicode/utf8"
)

func getArticleTagsInfo(id int) ([]message.TagInfo, error) {
	tags := make([]dao.Tag, 0)
	tagsInfoList := make([]message.TagInfo, 0)
	if err := dao.QueryAllTagsByArticleID(id, &tags); err != nil {
		msg := fmt.Sprintf("get all tags by article id fail, %v", id)
		juneLog.ExceptionLog(err, msg)
		return tagsInfoList, err
	}

	for _, tagInfo := range tags {
		tagsInfoList = append(tagsInfoList, message.TagInfo{
			ID:           tagInfo.ID,
			Name:         tagInfo.Name,
			CreateTime:   tagInfo.CreateTime.Unix(),
			ArticleTotal: tagInfo.Total,
		})
	}
	return tagsInfoList, nil
}

func ArticleTagsLogic(ctx *gin.Context,
	req juneGin.BaseReqInter) juneGin.BaseRespInter {
	reqL := req.(*message.ArticleTagsReq)
	resp := message.ArticleTagsResp{}
	tags, err := getArticleTagsInfo(reqL.ArticleID)
	if err != nil {
		mes := fmt.Sprintf("get article tags fail, "+
			"article id = %v ", reqL.ArticleID)
		juneLog.ExceptionLog(err, mes)
		return juneGin.SystemErrorRespHeader
	}
	resp.ID = reqL.ArticleID
	resp.Tags = tags
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func articleListByTag(tagID, page, pageSize int) (*message.ArticleListResp, error) {

	articleList, total, err := dao.QueryArticleInfoByLimitWithTag(tagID, page, pageSize)
	if err != nil {
		return nil, err
	}
	return &message.ArticleListResp{
		Header:      juneGin.SuccessRespHeader,
		ArticleList: articleList,
		Total:       total,
	}, nil
}

// 文章列表逻辑
func ArticleListLogic(ctx *gin.Context,
	req juneGin.BaseReqInter) juneGin.BaseRespInter {
	reqL := req.(*message.ArticleListReq)
	resp := message.ArticleListResp{}

	if reqL.Tag != 0 {
		response, err := articleListByTag(reqL.Tag, reqL.Page, reqL.PageSize)
		if err != nil {
			return juneGin.SystemErrorRespHeader
		}
		return response
	}

	articleList, total, err := dao.QueryArticleInfoByLimit(reqL.Page, reqL.PageSize)
	if err != nil {
		return juneGin.SystemErrorRespHeader
	}
	resp.ArticleList = articleList
	resp.Header = juneGin.SuccessRespHeader
	resp.Total = total
	return resp
}

func ArticleDetailLogic(ctx *gin.Context,
	req juneGin.BaseReqInter) juneGin.BaseRespInter {
	resp := message.ArticleDetailResp{}
	reqD := req.(*message.ArticleDetailReq)

	article, _ := dao.QueryArticleDetail(reqD.ArticleID)
	resp.ID = article.ID
	resp.Text = article.Text
	resp.CreateTime = article.CreateTime
	resp.Abstract = article.Abstract
	resp.AuthorID = article.AuthorID
	resp.Title = article.Title
	resp.BaseRespHeader = juneGin.SuccessRespHeader
	return resp
}

func ArticleAddLogic(ctx *gin.Context,
	req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.ArticleAddReq)
	resp := message.ArticleAddResp{}
	user, ok := ctx.Get("user")
	if !ok {
		resp.Header = juneGin.UnauthorizedRespHeader
		return resp
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
		resp.Header = juneGin.SystemErrorRespHeader
		return resp
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func getAbstract(text string) string {
	abstractList := strings.Split(text, consts.AbstractSplitStr)
	sp := src.GetSetting().Others.AbstractLen
	// 没有显示定义摘要，提取文字前部分内容作为摘要
	if len(abstractList) < 2 {
		if utf8.RuneCountInString(text) > sp {
			str := string([]rune(text)[:sp]) + "..."
			str = util.RemoveTitle(str)
			return strings.Replace(str, "\n", "", -1)
		}
		// 文章很短的情况
		str := util.RemoveTitle(text)
		return strings.Replace(str, "\n", "", -1)
	}

	r := util.RemoveTitle(abstractList[0])
	r = strings.Replace(r, "\n", "", len(r))
	if utf8.RuneCountInString(r) > sp {
		return string([]rune(r)[:sp-3]) + "..."
	}
	return r
}

func ArticleUpdateLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.ArticleUpdateReq)
	resp := message.ArticleUpdateResp{}

	user, ok := ctx.Get("user")
	if !ok {
		return juneGin.UnauthorizedRespHeader
	}
	author := user.(*dao.User)

	// if request.CreateTime.Unix() < 0 {
	// 	request.CreateTime = time.Now()
	// }

	abstract := request.Abstract
	if abstract == "" {
		abstract = getAbstract(request.Text)
	}

	// update article table
	err := dao.UpdateArticle(request.ID, &dao.Article{
		ID:       request.ID,
		Text:     request.Text,
		Title:    request.Title,
		AuthorID: author.ID,
		Abstract: abstract,
		// CreateTime: request.CreateTime,
	})
	if err != nil {
		return juneGin.SystemErrorRespHeader
	}
	// update article_tag table
	err = dao.UpdateArticleTagsByIntList(request.ID, request.Tags)
	if err != nil {
		return juneGin.SystemErrorRespHeader
	}

	resp.Header = juneGin.SuccessRespHeader
	return resp
}

func ArticleDeleteLogic(ctx *gin.Context, req juneGin.BaseReqInter) juneGin.BaseRespInter {
	request := req.(*message.ArticleDeleteReq)
	resp := message.ArticleDeleteResp{}
	err := dao.DeleteArticle(request.ID)
	if err != nil {
		return juneGin.SystemErrorRespHeader
	}
	resp.Header = juneGin.SuccessRespHeader
	return resp
}
