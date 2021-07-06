package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// GetSQLXMysqlInstance returns pointer to sqlx.DB
func GetSQLXMysqlInstance(
	username string,
	password string,
	host string,
	dbName string) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		dbName,
	)

	return sqlx.Connect("mysql", dsn)
}