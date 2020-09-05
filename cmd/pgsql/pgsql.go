package pgsql

import (
	"log"

	_ "github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"

	"daily/cmd/config"
)

var DB *sqlx.DB

func init() {
	var (
		err      error
		postgres = config.GetPostgres()
	)

	DB, err = sqlx.Connect(postgres.Driver, postgres.Host)
	if err != nil {
		log.Panic(err)
	}
	log.Println("success init postgreSQL")
}
