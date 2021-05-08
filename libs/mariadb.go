package libs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func MariaDBInit() *sqlx.DB {
	db := sqlx.MustConnect("mysql", "golanguser:golanguser@tcp(localhost:3306)/golang?charset=latin1")
	return db
}
