package task

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

/// usecaseをinjectionしたい場合は？

func Index(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	tmpl, err := template.ParseFiles("view/task/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Show(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")
	fmt.Println(taskId)

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task show: "+taskId)

}

func Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task store")
}

func Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task update")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task delete")
}
