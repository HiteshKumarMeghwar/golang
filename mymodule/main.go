package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page<p>To get in touch, email me at <a href=\"/\">Home page</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page!</h1>
		<ul>
			<li><b>Is here a free version?</b> Yes! We offer a free trail for 30 day on way paid plan.</li>
			<li><b>What are your support hours?</b> Yes! We offer a free trail for 30 day on way paid plan.</li>
			<li><b>How do I contact support?</b> Email  us - <a href="www.google.com">Google</a></li>
		</ul>
	`)
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
type Router struct{}

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

func main() {
	var router Router
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", router)
}
