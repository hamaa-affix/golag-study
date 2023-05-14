package task

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

/// usecaseをinjectionしたい場合は？

type TaskContorller struct {
}

func (t *TaskContorller) Index(w http.ResponseWriter, r *http.Request) {
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

func (t *TaskContorller) Show(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")
	fmt.Println(taskId)

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task show: "+taskId)

}

func (t *TaskContorller) Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task store")
}

func (t *TaskContorller) Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task update")
}

func (t *TaskContorller) Delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "task delete")
}
