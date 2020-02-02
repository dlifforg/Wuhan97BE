package main

import (
	"github.com/gin-gonic/gin"
	"hello/handler"
	"hello/service"
)

func main() {
	service.Init()
	r := gin.Default()
	r.GET("/index", handler.Index)
	service.Init()
	service.Load()
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}