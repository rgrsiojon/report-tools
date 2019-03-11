package main

import (
	"github.com/rgrsiojon/report-tools/config"

	"github.com/gin-gonic/gin"
	"github.com/rgrsiojon/report-tools/routers"
)

func main() {
	Config := config.ReadConfig()
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	r.Run(":" + Config.Server.Port)
}
