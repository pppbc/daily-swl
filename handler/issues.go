package handler

import (
	"errors"
	"github.com/gin-gonic/gin"

	"daily/service/issues"
	"daily/utils/apires"
	"daily/utils/apires/apierr"
)

func IssuesCreate(ctx *gin.Context) {
	var input issues.IssueInput
	if err := ctx.ShouldBind(&input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	//input.Time = time.Now().Format("2006-01-02")
	if err := issues.IssueController.Create(input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithNil(ctx)
	return
}

func IssuesUpdate(ctx *gin.Context) {
	var input issues.IssueInput
	if err := ctx.ShouldBind(&input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	if input.FinishIf {
		apierr.HandleErr(ctx, errors.New("已完成的任务不支持修改"))
		return
	}
	if err := issues.IssueController.Update(input.Id, input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithNil(ctx)
	return
}

func IssuesDelete(ctx *gin.Context) {
	var input issues.IssueInput
	if err := ctx.BindJSON(&input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	if err := issues.IssueController.Delete(input.UserId, input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithNil(ctx)
	return
}

func IssuesList(ctx *gin.Context) {
	var input issues.IssueParam
	if err := ctx.BindQuery(&input); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	var output []*issues.IssueOutput
	if err := issues.IssueController.List(input.UserId, input, &output); err != nil {
		apierr.HandleErr(ctx, err)
		return
	}
	apires.ResWithData(ctx, output)
	return
}

func IssuesDetail(ctx *gin.Context) {

}
