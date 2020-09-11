package router

import (
	"github.com/gin-gonic/gin"

	"daily/cmd/config"
	"daily/handler"
	"daily/utils/middleware"
)

func InitRouter() {
	router := gin.Default()

	router.Use(middleware.XOptions)
	router.GET("/daily/login", handler.Login)

	// token校验
	rt := router.Group("/daily", middleware.MiddlewareImpl)

	// 用户相关接口
	uGroup := rt.Group("/users")
	{
		uGroup.GET("", handler.UsersDetail)
		uGroup.POST("", handler.UsersUpdate)
		uGroup.PUT("/avatar", handler.UploadAvatar)
	}

	// 每日事项
	iGroup := rt.Group("/issues")
	{
		//iGroup.GET("", handler.IssuesDetail)
		iGroup.POST("", handler.IssuesCreate)
		iGroup.PUT("", handler.IssuesUpdate)
		iGroup.DELETE("", handler.IssuesDelete)
		iGroup.GET("/list", handler.IssuesList)
	}

	err := router.Run(config.GetIP())
	if err != nil {
		panic(err)
	}
}
