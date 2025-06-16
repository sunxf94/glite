package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg Config) (err error) {
	s := cfg.getDataSourceName()
	if db, err = sqlx.Open("mysql", s); err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return err
}

func InitWithSQLite(cfg Config) (err error) {
	if db, err = sqlx.Open("sqlite", cfg.DBName); err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return err
}

func GetDB() *sqlx.DB {
	return db
}
