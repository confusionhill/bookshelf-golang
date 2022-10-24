package database

import (
	"database/sql"
	"time"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:dukun123@(localhost:3306)/belajar")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
