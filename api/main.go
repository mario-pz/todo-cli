package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectDB()
	defer db.Close()

	r := chi.NewRouter()

	r.Post("/register", registerHandler(db))
	r.Post("/login", loginHandler(db))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
