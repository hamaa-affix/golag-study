package todo

import (
	"fmt"
	"net/http"
)
/// usecaseをinjectionしたい場合は？

func Index(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "todo index")
}

func Show(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "todo show")
}

func Store(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "todo store")
}

func Update(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "todo update")
}

func Delete(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.Form)
	fmt.Fprintf(w, "todo delete")
}
