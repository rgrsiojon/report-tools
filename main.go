package main

import (
	"./config"
	"./routers"
	"github.com/gin-gonic/gin"
)

func main() {
	Config := config.ReadConfig()
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	r.Run(":" + Config.Server.Port)
}
