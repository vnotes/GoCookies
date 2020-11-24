package mysqldb

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Pool *sqlx.DB
)

func init() {
	Pool = sqlx.MustOpen("mysql", "root:root@tcp(127.0.0.1:3306)/dev?parseTime=true")
	Pool.SetMaxOpenConns(1000)
	Pool.SetMaxIdleConns(10)

}
