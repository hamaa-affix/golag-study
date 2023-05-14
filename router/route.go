package router

import (
	// "fmt"
	"net/http"

	"golag-study/controller/task"

	"github.com/go-chi/chi/v5"
)

func Do() {
	r := chi.NewRouter()

	r.Route("/task", func(r chi.Router) {
		r.Get("/", task.Index)
		r.Get("/{taskId}", task.Show)
		r.Put("/{taskId}/edit", task.Update)
		r.Post("/store", task.Store)
		r.Delete("/{taskId}/delete", task.Delete)
	})

	http.ListenAndServe(":8080", r)
}
