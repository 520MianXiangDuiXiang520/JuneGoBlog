package check

import (
	"JuneGoBlog/internal/api/message"
	"errors"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	"github.com/gin-gonic/gin"
	"log"
)

func TagListCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	return nil, nil
}

func TagAddCheck(ctx *gin.Context, req juneGin.BaseReqInter) (juneGin.BaseRespInter, error) {
	reqA := req.(*message.TagAddReq)
	if reqA.TagName == "" {
		log.Printf("Check Add Tag Error! name = [%v]\n", reqA.TagName)
		return juneGin.ParamErrorRespHeader, errors.New("")
	}
	return nil, nil
}
