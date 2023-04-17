package router_demo

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func boot() *chi.Mux {
	r := chi.NewRouter()
	fmt.Println(r)
	routeRegister(r)
	return r
}

func routeRegister(router *chi.Mux) {
	fmt.Println(router)
	r := chi.NewRouter()

	// todo用のhandler
	r.Get("/todo", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hoge")
	}))

}
