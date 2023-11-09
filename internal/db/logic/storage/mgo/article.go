package mgo

import (
	"JuneGoBlog/internal/db/module"
	"JuneGoBlog/internal/db/opt"
)

func (m *Mgo) FindTotalCount(opts ...opt.Opt) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) FindSomeArticleInfo(opts ...opt.Opt) ([]module.ArticleHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) FindOneArticleInfo(id int64, opts ...opt.Opt) (module.ArticleHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) FindOneArticleDetail(id int64, opts ...opt.Opt) (module.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) HasArticle(id int, opts ...opt.Opt) bool {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) NewArticle(artifact *module.Article, opts ...opt.Opt) (*module.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) UpdateArticle(id int, article *module.Article, opts ...opt.Opt) error {
	//TODO implement me
	panic("implement me")
}

func (m *Mgo) DeleteArticle(id int, opts ...opt.Opt) error {
	//TODO implement me
	panic("implement me")
}
