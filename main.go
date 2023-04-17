package main

import (
	// "fmt"
	// "net/http"

	"golag-study/router"
)

func main() {
	// http.Handle("/todo", http.HandlerFunc(router.TodoIndexHandler))
	// http.Handle("/todo/post", http.HandlerFunc(router.TodoPostHandler))
	// http.Handle("/todo/1/update", http.HandlerFunc(router.TodoUpdateHandler))
	// http.Handle("/todo/1/delete", http.HandlerFunc(router.TodoDeleteHandler))

	// http.ListenAndServe(":8080", nil)
	// h1 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// }
	// h2 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #2!\n")
	// }

	// http.HandleFunc("/", h1)
	// http.HandleFunc("/endpoint", h2)

	// log.Fatal(http.ListenAndServe(":8080", nil))
	router.Do()
}
