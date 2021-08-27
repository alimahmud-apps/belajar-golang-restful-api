package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

const (
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "belajar_golang_restful_api"
)

var host string = os.Getenv("MYSQL_HOST")

func NewDB() *sql.DB {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", mysqlInfo)

	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
