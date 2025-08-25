package main

import (
	"github.com/gin-gonic/gin"
	"server/core"
	"server/global"
	"server/middleware"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	router := gin.Default()
	router.Use(middleware.GinLogger(), middleware.GinRecover(true))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
