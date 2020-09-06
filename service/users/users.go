package users

import (
	"github.com/jmoiron/sqlx"

	"daily/cmd/pgsql"
)

type userController struct {
	db *sqlx.DB
}

var UserController = &userController{db: pgsql.DB}

func (u *userController) Get(id int64, output interface{}) error {
	query := `SELECT * FROM daily.users WHERE id = $1`
	return u.db.Get(output, query, id)
}

func (u *userController) Update(id int64, input UserInput, output interface{}) error {
	tx := u.db.MustBegin()

	query := `UPDATE daily.users SET (name,sex) = ($1,$2) WHERE id = $3 RETURNING *;`
	if err := tx.Get(output, query, input.Name, input.Sex, id); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (u *userController) UpdateLoginTime(id int64, output interface{}) error {
	tx := u.db.MustBegin()

	query := `UPDATE daily.users SET login_at = current_timestamp WHERE id = $1 RETURNING *;`
	if err := tx.Get(output, query, id); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (u *userController) FindByName(input LoginParams, output interface{}) error {
	query := `SELECT * FROM daily.users WHERE name = $1;`
	return u.db.Get(output, query, input.Username)
}

func (u *userController) UpdateAvatar(id int64, path string, output interface{}) error {
	tx := u.db.MustBegin()

	query := `UPDATE daily.users SET avatar = $1 WHERE id = $2 RETURNING *;`
	if err := tx.Get(output, query, path, id); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (u *userController) ChangePassword(id int64, newPwd string) error {
	tx := u.db.MustBegin()

	query := `UPDATE daily.users SET password = $1 WHERE id = $2;`
	if _, err := tx.Exec(query, newPwd, id); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
