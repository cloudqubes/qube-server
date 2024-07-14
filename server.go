package main

import (
	"fmt"
	"log"
	"net/http"
	// "text/template"
)

var count int = 0

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cloud\n")
	log.Println(r.Method, r.URL)
	// t, _ := template.ParseFiles("views/home.html")
	// t.Execute(w, "")
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	count += 1
	fmt.Fprintf(w, "%s%d%s", "Count: ", count, "\n")
	log.Println(r.Method, r.URL)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/count", countHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
