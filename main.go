package JuneGoBlog

import (
	"github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()
    defer engine.Run()
    Routes(engine)
}
