package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

// func about(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Contact Page<p>To get in touch, email me at <a href=\"www.google.com\">Google</a>.</p>")
// }

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", nil)
}
