package db

type IDbOperation interface {
	Init()
	IArticleDbOperation
}

var _db IDbOperation

func RegisterDbOp(op IDbOperation) {
	_db = op
}

func Db() IDbOperation {
	return _db
}

type IDbConfig interface {
	GetURI() string
	Timeout() int64 // ms
	PoolMaxSize() int
	PoolMinSize() int
}

func InitDb() {
	Db().Init()
}
