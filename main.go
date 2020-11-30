package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleSlash)
	http.ListenAndServe(":8080", nil)
}

func handleSlash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "<h1>My name is Rohit</h1>")
}
