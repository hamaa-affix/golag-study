package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golag-study/controller/todo"
)

func Do() {
	r := chi.NewRouter()
	r.Route("/todo", func(r chi.Router) {
		r.Get("/", todo.Index)
		r.Get("/1", todo.Show)
		r.Put("/1/edit", todo.Update)
		r.Post("/store", todo.Store)
		r.Delete("/1/delete", todo.Delete)
	})
	
	http.ListenAndServe(":8080", r)
}
