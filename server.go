package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cloud")
	t, _ := template.ParseFiles("views/home.html")
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
