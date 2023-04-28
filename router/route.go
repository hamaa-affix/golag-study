package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golag-study/controller/task"
)

func Do() {
	r := chi.NewRouter()
	r.Route("/task", func(r chi.Router) {
		r.Get("/", task.Index)
		r.Get("/1", task.Show)
		r.Put("/1/edit", task.Update)
		r.Post("/store", task.Store)
		r.Delete("/1/delete", task.Delete)
	})

	http.ListenAndServe(":8080", r)
}
