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

func main() {
	mux := http.NewServeMux()
	html = template.Must(template.ParseFiles("templates/index.html"))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
