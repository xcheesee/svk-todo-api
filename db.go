package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	var db *sql.DB
	var err error

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "svk_todo",
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return db
}
