package mgo

import (
	"JuneGoBlog/internal/config"
	"JuneGoBlog/internal/db"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var _ db.IDbOperation = (*Mgo)(nil)

type Mgo struct {
	cli *mongo.Client
}

func NewMgo() *Mgo {
	return &Mgo{}
}

func (m *Mgo) newConn() {
	c, ok := config.Cfg().GetVal("db.mongo")
	if !ok {
		log.Panicln("DB: mongo config load fail!")
	}
	cfg, ok := c.(*config.MongoCfg)
	if !ok {
		log.Panicln("DB: not mongo config, check unmarshal!")
	}

	opts := options.Client()
	opts.ApplyURI(cfg.URI)
	opts.SetCompressors(cfg.Compressors)
	//opts.SetConnectTimeout(time.Duration(cfg.ConnTimeout) * time.Second)
	opts.SetMaxConnecting(cfg.MaxConnecting)
	opts.SetMaxPoolSize(cfg.MaxPoolSize)
	opts.SetMinPoolSize(cfg.MinPoolSize)
	ctx, done := context.WithTimeout(context.Background(),
		time.Duration(cfg.ConnTimeout)*time.Second)
	defer done()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Panicf("connect mongo fail: addr: %s, err %s", cfg.URI, err.Error())
	}
	m.cli = client
}

func (m *Mgo) Init() {
	m.newConn()
}
