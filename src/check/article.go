package check

import (
	"JuneGoBlog/src/dao"
	junebaotop "JuneGoBlog/src/junebao.top"
	"JuneGoBlog/src/message"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleListCheck(ctx *gin.Context, req junebaotop.BaseReqInter) (junebaotop.BaseRespInter, error) {
	reqL := req.(*message.ArticleListReq)
	if reqL.PageSize <= 0 || reqL.Page <= 0 {
		return junebaotop.ParamErrorRespHeader, errors.New("ParamError")
	}
	if reqL.Tag != 0 {
		if _, ok := dao.HasTagByID(reqL.Tag); !ok {
			return junebaotop.ParamErrorRespHeader, errors.New("TagNotFind")
		}
	}
	return http.StatusOK, nil
}

func ArticleDetailCheck(ctx *gin.Context, req junebaotop.BaseReqInter) (junebaotop.BaseRespInter, error) {
	reqD := req.(*message.ArticleDetailReq)
	if reqD.ArticleID == 0 {
		return junebaotop.ParamErrorRespHeader, errors.New("ParamError")
	}
	return http.StatusOK, nil
}

func ArticleTagsCheck(ctx *gin.Context, req junebaotop.BaseReqInter) (junebaotop.BaseRespInter, error) {
	reqL := req.(*message.ArticleTagsReq)
	if reqL.ArticleID == 0 {
		return junebaotop.ParamErrorRespHeader, errors.New("ParamError")
	}
	return http.StatusOK, nil
}

func ArticleAddCheck(ctx *gin.Context, req junebaotop.BaseReqInter) (junebaotop.BaseRespInter, error) {
	request := req.(*message.ArticleAddReq)
	if len(request.Title) == 0 || len(request.Text) == 0 ||
		len(request.Tags) == 0 {
		return junebaotop.ParamErrorRespHeader, errors.New("")
	}
	return http.StatusOK, nil
}

func ArticleUpdateCheck(ctx *gin.Context, req junebaotop.BaseReqInter) (junebaotop.BaseRespInter, error) {
	request := req.(*message.ArticleUpdateReq)
	if len(request.Title) == 0 || len(request.Text) == 0 ||
		len(request.Tags) == 0 || !dao.HasArticle(request.ID) {
		return junebaotop.ParamErrorRespHeader, errors.New("")
	}
	return http.StatusOK, nil
}
