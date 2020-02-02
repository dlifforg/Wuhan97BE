package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"hello/handler"
	"hello/service"
	"log"
)

func main() {
	service.Init()
	r := gin.Default()
	r.GET("/index", handler.Index)
	service.Init()

	c := cron.New()
	c.AddFunc("0 */5 * * * *", func() {
		log.Print("starting cron task ")
		service.Load()
		log.Print("cron task finished")
	})
	c.Start()

	r.Run() // 在 0.0.0.0:8080 上监听并服务
}