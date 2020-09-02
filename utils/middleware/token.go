package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"daily/utils/safes"
)

func MiddlewareImpl(ctx *gin.Context) {
	authorization := ""
	authorization = ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		authorization = "Bearer " + ctx.Query("token")
	}

	token, err := safes.Get(authorization)
	if err != nil {
		err := errors.New("token error:" + err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": err.Error()})
		ctx.Abort()
		return
	}

	claimsMap, err := safes.Parse(token)
	if err != nil {
		err := errors.New("token error:" + err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": err.Error()})
		ctx.Abort()
		return
	}
	//role := claimsMap["role"].(string)
	userid := claimsMap["userid"].(string)
	//if singlePoint == true {
	//	lt := claimsMap["logintime"].(string)
	//	logintime, err := controller.Middleware(role, userid)
	//	if err != nil {
	//		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": err.Error()})
	//		ctx.Abort()
	//		return
	//	}
	//	if strconv.FormatInt(logintime, 10) != lt {
	//		err := errors.New("您的帐号已在其他地方登录")
	//		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": err.Error()})
	//		ctx.Abort()
	//		return
	//	}
	//}

	//ctx.Set("ROLE", role)
	ctx.Set("USERID", userid)
	ctx.Next()
}
