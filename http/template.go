package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Page{"Home", "Welcome to Go!"}
		tmpl.Execute(w, page)
	})
	http.ListenAndServe(":8080", nil)
}
