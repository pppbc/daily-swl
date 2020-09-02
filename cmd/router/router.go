package router

import (
	"github.com/gin-gonic/gin"

	"daily/handler"
	"daily/utils/middleware"
)

func InitRouter() {
	router := gin.Default()

	router.POST("/daily/login", handler.Login)

	// token校验
	rt := router.Group("/daily", middleware.MiddlewareImpl)

	// 用户相关接口
	uGroup := rt.Group("/users")
	{
		uGroup.GET("", handler.UsersDetail)
		uGroup.POST("", handler.UsersUpdate)
	}

	// 每日事项
	iGroup := rt.Group("/issues")
	{
		iGroup.GET("", handler.IssuesDetail)
		iGroup.POST("", handler.IssuesCreate)
		iGroup.PUT("", handler.IssuesUpdate)
		iGroup.DELETE("", handler.IssuesDelete)
		iGroup.GET("/list", handler.IssuesList)
	}

}
