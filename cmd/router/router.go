package router

import (
	"daily/cmd/config"
	"github.com/gin-gonic/gin"

	"daily/handler"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/daily/login", handler.Login)

	// token校验
	//rt := router.Group("/daily", middleware.MiddlewareImpl)
	rt := router.Group("/daily")

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
