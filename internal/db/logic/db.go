package logic

import (
	"JuneGoBlog/internal/db"
	"JuneGoBlog/internal/db/logic/storage/mgo"
	"JuneGoBlog/internal/db/module"
	"JuneGoBlog/internal/db/opt"
)

type ILogic interface {
	Cache() db.IDbOperation
	Storage() db.IDbOperation
}

var (
	_ db.IDbOperation = (*dbLogic)(nil)
	_ ILogic          = (*dbLogic)(nil)
)

func init() {
	db.RegisterDbOp(newDbLogic())
}

type dbLogic struct {
	c db.IDbOperation
	s db.IDbOperation
}

func newDbLogic() *dbLogic {
	d := &dbLogic{}
	d.s = mgo.NewMgo()
	d.c = nil
	return d
}

func (d *dbLogic) Cache() db.IDbOperation {
	return d.c
}

func (d *dbLogic) Storage() db.IDbOperation {
	return d.s
}

func (d *dbLogic) Init() {
	d.Storage().Init()
	d.Cache().Init()
}

func (d *dbLogic) FindTotalCount(opts ...opt.Opt) (int, error) {
	if n, err := d.Cache().FindTotalCount(opts...); err == nil {
		return n, nil
	}
	// cache miss
	return d.Storage().FindTotalCount(opts...)
}

func (d *dbLogic) FindSomeArticleInfo(opts ...opt.Opt) ([]module.ArticleHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) FindOneArticleInfo(id int64, opts ...opt.Opt) (module.ArticleHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) FindOneArticleDetail(id int64, opts ...opt.Opt) (module.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) HasArticle(id int, opts ...opt.Opt) bool {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) NewArticle(artifact *module.Article, opts ...opt.Opt) (*module.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) UpdateArticle(id int, article *module.Article, opts ...opt.Opt) error {
	//TODO implement me
	panic("implement me")
}

func (d *dbLogic) DeleteArticle(id int, opts ...opt.Opt) error {
	//TODO implement me
	panic("implement me")
}
