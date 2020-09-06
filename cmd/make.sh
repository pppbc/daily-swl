#!/bin/sh
###生成api文档
#swag init

###关闭C语言版本的编译器
export CGO_ENABLED=0

###格式化、标准化golang代码
gofmt -w *.go && gofmt -w */*.go

export GO111MODULE=on
export GOPROXY=https://goproxy.io
export GOOS=linux
export GOARCH=amd64

###定义项目可执行程序名称和服务入口main文件
PROJECT_NAME=daily
PROJECT_MAIN=main.go

# set project version
GIT_COMMIT_VERSION=`date -u +%Y%m%d.%H%M%S`
if git status >/dev/null 2>&1 ; then 
    GIT_COMMIT_VERSION=r`git rev-parse --short HEAD`
fi

### 编译项目
#GO_LDFLAGS="$GO_LDFLAGS -X dsight/common/version.commitVersion=$GIT_COMMIT_VERSION"
go build -ldflags "$GO_LDFLAGS" -o $PROJECT_NAME $PROJECT_MAIN
