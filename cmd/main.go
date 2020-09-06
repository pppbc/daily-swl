package main

import (
	"daily/cmd/logger"
	"daily/cmd/router"
)

func main() {
	logger.LogInitAndStart()
	// 初始化路由
	router.InitRouter()
	//timeStr := time.Now().Format("2006-01-02")

	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(timeStr)
}
