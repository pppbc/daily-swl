package users

// 用户表模型
type UserModels struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Avatar   string `json:"avatar" db:"avatar"`
	Password string `json:"password" db:"password"`
	Sex      int    `json:"sex" db:"sex"`
	CreateAt string `json:"create_at" db:"create_at"`
	LoginAt  string `json:"login_at" db:"login_at"`
}

// 登录参数
type LoginParams struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 修改信息
type UserInput struct {
	Id       int64  `json:"id" db:"id" form:"id"`
	Name     string `json:"name" db:"name" form:"name"`
	Avatar   string `json:"avatar" db:"avatar"`
	Password string `json:"password" db:"password" form:"password"`
	Sex      int    `json:"sex" db:"sex" form:"sex"`
	CreateAt string `json:"create_at" db:"create_at" form:"create_at"`
	LoginAt  string `json:"login_at" db:"login_at" form:"login_at"`
}

// 修改信息
type UserOutput struct {
	Id       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Avatar   string  `json:"avatar" db:"avatar"`
	Password string  `json:"password" db:"password"`
	Sex      int     `json:"sex" db:"sex"`
	CreateAt string  `json:"create_at" db:"create_at"`
	LoginAt  *string `json:"login_at" db:"login_at"`
}
