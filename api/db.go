package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// TODO: This will be replaced with environment variables
	var driver = "mysql"
	var user = "user"
	var password = "bingus"
	var database = "tododb"
	var databaseHost = "0000:3306"

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
