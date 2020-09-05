package apires

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {
	Result interface{} `json:"result"`
	Total  int64       `json:"total"`
}
type ResMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	// 206 成功执行部分请求
	PasswordIncorrect = ResMessage{206, "password incorrect"}
	UsernameNotFound  = ResMessage{206, "no such user"}
	UserDeleted       = ResMessage{206, "user deleted"}
)
var (
	// not found
	TotalNotFound = ResMessage{206, "no result"}
)

// 做出响应
func ResWithData(r *gin.Context, data interface{}) {
	r.JSON(http.StatusOK, data)
}

func ResWithNil(r *gin.Context) {
	r.JSON(http.StatusOK, nil)
}

func ResWithMessage(r *gin.Context, message ResMessage) {
	r.JSON(message.Status, message.Message)
}
