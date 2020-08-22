package main

import (
	"JuneGoBlog/src"
	"github.com/gin-gonic/gin"
)

func init() {
	src.InitSetting("./setting.json")
}

func main() {
	engine := gin.Default()
	defer engine.Run()
	Register(engine)
}
