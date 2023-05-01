package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := connectDB()
	defer db.Close()

	r := chi.NewRouter()

	// Public Routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", registerHandler(db))
		r.Post("/login", loginHandler(db))

		r.Post("/create_task", createTask(db))

		r.Delete("/delete_task", deleteTask(db))

		r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World"))
		})
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
