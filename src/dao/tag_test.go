package dao

import (
	"JuneGoBlog/src"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	"testing"
)

func init() {
	src.InitSetting("../../setting.json")
	setting := src.GetSetting()
	// 初始化数据库连接
	_ = juneDao.InitDBSetting(setting.MySQLSetting)
}

func TestQueryArticleTotalByTagIDFromCache(t *testing.T) {

}

func TestQueryTagByIDFromCache(t *testing.T) {
	_, _ = queryTagByIDFromCache(13)
}

func TestQueryTagArticleTotal(t *testing.T) {
	total, err := QueryTagArticleTotal(juneDao.GetDB(), 10)
	if err != nil || total != 1 {
		t.Error()
	}
}
