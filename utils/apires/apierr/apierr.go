package apierr

import (
	"github.com/gin-gonic/gin"
)

type ErrHandler struct {
	Status int    `json:"status"`
	Err    string `json:"err"`
}

type Err struct {
	Status int   `json:"status"`
	Err    error `json:"err"`
}

// 自定义逻辑
var (
	ErrNotFound  = ErrHandler{404, "page not found"}
	ErrBindBody  = ErrHandler{400, "request body bind failed"}
	ErrBindForm  = ErrHandler{400, "request form bind failed"}
	ErrBindParam = ErrHandler{400, "request param bind failed"}
	NoUserID     = ErrHandler{400, "no user id"}
	RoleIsNull   = ErrHandler{400, "role is null"}
	InputIsNull  = ErrHandler{400, "input is null"}
	NoAccess     = ErrHandler{400, "no access"}
)

// user
var (
	UsernameIsNull         = ErrHandler{400, "username is null"}
	PasswordStandardIsNull = ErrHandler{400, "password standard is null"}
	EmailIsInUse           = ErrHandler{400, "email is in use"}
	PhoneIsInUse           = ErrHandler{400, "phone is in use"}
	EmailCanNotFind        = ErrHandler{400, "email can not find"}
	PhoneCanNotFind        = ErrHandler{400, "phone can not find"}
	ResultNotFound         = ErrHandler{400, "result not found"}
	CodeNotrueT            = ErrHandler{400, "wrong code"}
)
var (
	PasswordIncorrect = ErrHandler{400, "password incorrect"}
	UsernameNotFound  = ErrHandler{400, "no such user"}
	UserDeleted       = ErrHandler{400, "user deleted"}
)

// photo

var (
	PhotoNotFound = ErrHandler{400, "photo not found"}
)

// service

var (
	ServiceNotFound = ErrHandler{400, "service not found"}
)

func HandlerOwnErr(r *gin.Context, res ErrHandler) {
	r.JSON(res.Status, res.Err)
}

func HandleErr(r *gin.Context, err error) {
	// Err 返回不了,why,太长了??????
	r.JSON(400, err)
}
