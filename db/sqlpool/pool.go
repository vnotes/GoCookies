package mysql

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbPool *sqlx.DB
)

func init() {
	dbPool = sqlx.MustOpen("mysql", "root:root@tcp(127.0.0.1:3306)/dev")
	dbPool.SetMaxOpenConns(1000)
	dbPool.SetMaxIdleConns(10)

}
