package main

import (
	"log"

	"daily/cmd/logger"
	"daily/cmd/router"
)

func main() {
	err := logger.LogInitAndStart()
	if err != nil {
		log.Println("[log] init err:", err)
		//return
	}
	// 初始化路由
	router.InitRouter()
}
