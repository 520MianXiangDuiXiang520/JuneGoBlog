package dao

import (
	"JuneGoBlog/src/consts"
	"log"
)

// 用来预热缓存
var articleList []Article

func init() {
	articleList = make([]Article, 0)
	if err := QueryAllArticle(&articleList); err != nil {
		log.Println("InitArticleIDListCache Error!!")
		return
	}
}

// 文章列表预热
//   1. 预热 文章ID 列表
//   2. 预热 文章简单信息

func InitArticleIDListCache() error {
	// 1. 更新缓存中的 articleIDList
	rc := RedisPool.Get()
	defer rc.Close()
	for _, article := range articleList {
		_ = rc.Send("RPUSH", consts.ArticleIDListCache, article.ID)
	}
	_ = rc.Flush()
	return nil
}
