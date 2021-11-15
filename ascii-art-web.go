package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test1!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 2")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}