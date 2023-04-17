package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Do() {
	r := chi.NewRouter()

	// todoのルーティング
	r.Get("/todo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "todo")
	})
}
