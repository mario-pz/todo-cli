package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func connectDB() *sql.DB {
	//TODO: This will replaced with env variables
	var driver = "mysql"
	var user = "user"
	var password = "password"
	var database = "tododb"
	var databaseHost = "localhost:3306"

	var link = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, databaseHost, database)

	db, err := sql.Open(driver, link)

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func registerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(
			"INSERT INTO users (username, password) VALUES (?, ?)",
			user.Username,
			user.Password,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
	}
}

func loginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var password string
		err = db.QueryRow("SELECT password FROM users WHERE username=?", user.Username).
			Scan(&password)

		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		if password != user.Password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}
}
