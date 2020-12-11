package main

import (
	"github.com/gin-gonic/gin"
	"suiyidian.cn/sonhineboy/ginedu/controllers"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/ping", controllers.Test)
	r.GET("/index", controllers.Index)

	r.GET("/grsync", controllers.GoRsync)
	r.GET("/dblink", controllers.DbList)
	r.GET("/dblink2", controllers.DbGames)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
