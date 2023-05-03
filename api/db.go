package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	var driver = "mysql"
	var user = os.Getenv("MYSQL_USER")
	var password = os.Getenv("MYSQL_PASSWORD")
	var database = os.Getenv("MYSQL_DATABASE")
	var databaseHost = os.Getenv("MYSQL_HOST") + ":3306"

	var link = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, databaseHost, database)

	db, err := sql.Open(driver, link)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
