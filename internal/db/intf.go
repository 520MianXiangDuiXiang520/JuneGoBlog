package db

import (
	"JuneGoBlog/internal/db/module"
	"JuneGoBlog/internal/db/opt"
)

type IArticleDbOperation interface {
	FindTotalCount(opts ...opt.Opt) (int, error)

	FindSomeArticleInfo(opts ...opt.Opt) ([]module.ArticleHeader, error)

	FindOneArticleInfo(id int64, opts ...opt.Opt) (module.ArticleHeader, error)

	FindOneArticleDetail(id int64, opts ...opt.Opt) (module.Article, error)

	HasArticle(id int, opts ...opt.Opt) bool

	NewArticle(artifact *module.Article, opts ...opt.Opt) (*module.Article, error)

	UpdateArticle(id int, article *module.Article, opts ...opt.Opt) error

	DeleteArticle(id int, opts ...opt.Opt) error
}
