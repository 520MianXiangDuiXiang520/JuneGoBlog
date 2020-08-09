package main

import (
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	err := dao.InitDB()
	if err != nil {
		return
	}
}

func main() {
	engine := gin.Default()
	defer engine.Run()
	Register(engine)
	defer dao.DB.Close()
}
