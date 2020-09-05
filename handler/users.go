package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"daily/cmd/config"
	"daily/service/users"
	"daily/utils/apires"
	"daily/utils/apires/apierr"
)

func Login(ctx *gin.Context) {
	var params users.LoginParams
	err := ctx.BindQuery(&params)
	if err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	var output users.UserOutput
	if err := users.UserController.FindByName(params, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	if output.Password != params.Password {
		apierr.HandleErr(ctx, errors.New("密码不正确"))
		return
	}
	// todo 生成token
	if err := users.UserController.UpdateLoginTime(output.Id, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithData(ctx, output)
}

func UsersDetail(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		apierr.HandleErr(ctx, err)
	}

	var output users.UserOutput
	if err := users.UserController.Get(userId, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithData(ctx, output)
}
func UsersUpdate(ctx *gin.Context) {
	var params users.UserInput
	err := ctx.BindJSON(&params)
	if err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	var output users.UserOutput
	if err := users.UserController.Update(params.Id, params, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithData(ctx, output)
}

func UploadAvatar(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	filename := fmt.Sprintf("%davatar%02v.png", userId, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100))
	fileServer := config.GetFileServer()

	if err := ctx.SaveUploadedFile(file, fileServer.PhotoPath+filename); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	var output users.UserOutput
	if err := users.UserController.UpdateAvatar(userId, fileServer.Photo+filename, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}

	apires.ResWithData(ctx, output)
	return
}
