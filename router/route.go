package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golag-study/controller/todo"
)

func Do() {
	r := chi.NewRouter()

	// todoのルーティング
	r.Group(func(r chi.Router) {
		r.Get("/todo", todo.Index)
	})
	// r.Get("/todo/1", func(rw http.ResponseWriter, r *http.Request) {
	// 	if err := r.ParseForm(); err != nil {
	// 		http.Error(rw, err.Error(), http.StatusBadRequest)
	// 	}

	// 	fmt.Println(r.Form)
	// 	fmt.Fprintf(rw, "todo")
	// })

	http.ListenAndServe(":8080", r)
}
