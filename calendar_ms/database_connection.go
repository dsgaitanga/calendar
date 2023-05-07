package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const AllowedCORSDomain = "172.17.0.2"

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", "calendar:2023@tcp("+AllowedCORSDomain+":3306)/calendar_db?parseTime=true")
}
