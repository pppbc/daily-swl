package issues

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"daily/cmd/pgsql"
)

type issueController struct {
	db *sqlx.DB
}

var IssueController = &issueController{db: pgsql.DB}

func (u *issueController) Count(output interface{}) error {
	query := `SELECT COUNT(*) FROM daily.issues WHERE 1 = 1 `
	return u.db.Get(output, query)
}

func (u *issueController) List(id int64, param IssueParam, output *[]*IssueOutput) error {
	query := `SELECT * FROM daily.issues WHERE time = $1 ORDER BY finish_if = false DESC ,user_id = $2 DESC `
	fmt.Print(param.UserId, id)
	return u.db.Select(output, query, param.Time, id)
}

func (u *issueController) Update(id int64, input IssueInput) error {
	tx := u.db.MustBegin()

	query := `UPDATE daily.issues SET (name,level,update_at) = ($1,$2,current_timestamp) WHERE id = $3 AND finish_if = false;`
	if _, err := tx.Exec(query, input.Name, input.Level, id); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (u *issueController) Delete(id int64, param IssueInput) error {
	tx := u.db.MustBegin()

	query := `DELETE FROM daily.issues WHERE id = $1 and user_id = $2;`
	if _, err := tx.Exec(query, id, param.UserId); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (u *issueController) Create(input IssueInput) error {
	tx := u.db.MustBegin()
	query1 := `
			INSERT INTO daily.issues 
				(name,user_id,level,time,finish_if,check_if,create_at,update_at) 
			VALUES 
				($1,$2,$3,$4,false,false,current_timestamp,current_timestamp) ;`
	if _, err := tx.Exec(query1, input.Name, input.UserId, input.Level, input.Time); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
