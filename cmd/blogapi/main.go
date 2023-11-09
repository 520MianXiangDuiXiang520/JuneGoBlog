package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	doInit()
}

func main() {
	engine := gin.Default()
	defer engine.Run()
	Register(engine)
}
