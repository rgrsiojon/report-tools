package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rgrsiojon/report-tools/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	r.Run(":8080")
}
