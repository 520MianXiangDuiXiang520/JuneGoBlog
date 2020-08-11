package main

import (
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
)



func main() {
	engine := gin.Default()
	defer engine.Run()
	Register(engine)
}
