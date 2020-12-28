package dao

import (
	"JuneGoBlog/src/consts"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	"log"
	"sync"
)

// 用来预热缓存
var articleList []Article
var articleListLock sync.Mutex

// 文章列表预热
//   1. 预热 文章ID 列表
//   2. 预热 文章简单信息

func InitArticleIDListCache() error {
	// 1. 更新缓存中的 articleIDList
	rc := juneDao.GetRedisConn()
	defer rc.Close()
	if len(articleList) == 0 {
		articleListLock.Lock()
		defer articleListLock.Unlock()
		if len(articleList) == 0 {
			articleList = make([]Article, 0)
			if err := QueryAllArticle(&articleList); err != nil {
				log.Println("InitArticleIDListCache Error!!")
				return err
			}
		}
	}
	for _, article := range articleList {
		_ = rc.Send("RPUSH", consts.ArticleIDListCache, article.ID)
	}
	_ = rc.Flush()
	return nil
}
