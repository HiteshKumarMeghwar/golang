package main

import (
	"log"
	"net/http"
	"text/template"
)

var html *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO LIST",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Like this vide", Done: false},
		},
	}

	html.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO LIST",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Like this vide", Done: false},
		},
	}

	html.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	html = template.Must(template.ParseFiles("templates/about.html"))
	html = template.Must(template.ParseFiles("templates/index.html"))
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/todo", todo)
	/* err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	} */
	log.Fatal(http.ListenAndServe(":8080", mux))
}
