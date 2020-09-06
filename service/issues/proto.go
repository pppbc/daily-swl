package issues

// 日程表模型
type IssueModels struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	UserId   int64  `json:"user_id" db:"user_id"`
	Level    int    `json:"level" db:"level"`
	Time     string `json:"time" db:"time"`
	FinishIf bool   `json:"finish_if" db:"finish_if"`
	CheckIf  bool   `json:"check_if" db:"check_if"`
	CheckId  int64  `json:"check_id" db:"check_id"`
	CreateAt string `json:"create_at" db:"create_at"`
	UpdateAt string `json:"update_at" db:"update_at"`
}

type IssueParam struct {
	Time   string `json:"time"  form:"time"`
	UserId int64  `json:"user_id"  form:"user_id"`
}

type IssueInput struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	UserId   int64  `json:"user_id" db:"user_id"`
	Level    int    `json:"level" db:"level"`
	Time     string `json:"time" db:"time"`
	FinishIf bool   `json:"finish_if" db:"finish_if"`
	CheckIf  bool   `json:"check_if" db:"check_if"`
	CheckId  int64  `json:"check_id" db:"check_id"`
	CreateAt string `json:"create_at" db:"create_at"`
	UpdateAt string `json:"update_at" db:"update_at"`
}

type IssueOutput struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	UserId   int64  `json:"user_id" db:"user_id"`
	Level    int    `json:"level" db:"level"`
	Time     string `json:"time" db:"time"`
	FinishIf bool   `json:"finish_if" db:"finish_if"`
	CheckIf  bool   `json:"check_if" db:"check_if"`
	CheckId  *int64 `json:"check_id" db:"check_id"`
	CreateAt string `json:"create_at" db:"create_at"`
	UpdateAt string `json:"update_at" db:"update_at"`
}
