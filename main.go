package main

import (
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
	"log"
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
	defer log.Println("end。。。。")
}
