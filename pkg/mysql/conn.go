package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToRDS() {
	var err error
	db, err = sql.Open("", "")

	if err != nil {
		panic(err.Error())
	}
}

func ConnClose() {
	db.Close()
}
