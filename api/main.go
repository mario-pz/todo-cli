package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := ConnectDB()
	defer db.Close()

	r := chi.NewRouter()

	// Public Routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", RegisterHandler(db))
		r.Post("/login", LoginHandler(db))
		r.Post("/create_task", CreateTask(db))
		r.Delete("/delete_task", DeleteTask(db))
		r.Get("/gather_tasks", GatherTasks(db))
		r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World"))
		})
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
