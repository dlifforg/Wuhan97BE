package main

import (
	"github.com/robfig/cron"
	"hello/service"
	"log"
)

func main() {
	log.Println("Starting...")

	service.Init()
	c := cron.New()
	c.AddFunc("* */10 * * * *", func() {
		service.Load()
	})
	c.Start()

	select {
	}

}
