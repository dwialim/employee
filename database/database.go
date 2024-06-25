package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	db, err := sql.Open("mysql", "remote:remote@tcp(localhost:3306)/go_app")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
