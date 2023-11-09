package check

import (
	"JuneGoBlog/internal/api/message"
	"JuneGoBlog/internal/consts"
	"JuneGoBlog/internal/db/old"
	"errors"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

func ArticleListCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqL := req.(*message.ArticleListReq)
	if reqL.Tag != 0 {
		if _, ok := old.HasTagByID(reqL.Tag); !ok {
			return juneGin.ParamErrorRespHeader, errors.New("TagNotFind")
		}
	}
	return http.StatusOK, nil
}

func ArticleDetailCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	return http.StatusOK, nil
}

func ArticleTagsCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	return http.StatusOK, nil
}

func ArticleAddCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.ArticleAddReq)
	errResp := message.ArticleAddResp{
		Header: juneGin.ParamErrorRespHeader,
	}
	if utf8.RuneCountInString(request.Title) > consts.MaxArticleTitleLen {
		return errResp, errors.New("TitleTooLong")
	}
	return http.StatusOK, nil
}

func ArticleUpdateCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.ArticleUpdateReq)
	if !old.HasArticle(request.ID) {
		return juneGin.ParamErrorRespHeader, errors.New("")
	}
	return http.StatusOK, nil
}

func ArticleDeleteCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	request := req.(*message.ArticleDeleteReq)
	if !old.HasArticle(request.ID) {
		return nil, errors.New("ArticleDoesNotExist")
	}
	return http.StatusOK, nil
}
