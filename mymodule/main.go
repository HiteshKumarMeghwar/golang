package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"mymodule/views"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := views.Parse(filepath)

	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, filepath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, filepath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, filepath)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// fmt.Fprint(w, r.URL.Path)
// fmt.Fprint(w, r.URL.Path)
// fmt.Fprint(w, r.URL.RawPath)

/* if r.URL.Path == "/" {
	homeHandler(w, r)
	return
} else if r.URL.Path == "/contact" {
	contactHandler(w, r)
	return
} */

/* switch r.URL.Path {
case "/":
	homeHandler(w, r)
case "/contact":
	contactHandler(w, r)
default:

	//w.WriteHeader(http.StatusNotFound)
	//fmt.Fprint(w, "Ops Page not found!")

	// http.Error(w, "Ops Page Not Found!", http.StatusNotFound)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
} */

// }

// another way router
/* type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
*/

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)

	// var router Router

	// 3rd way of routing with chi
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found!", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", r)
}
