package main

import (
	"fmt"
	"net/http"
)

func main() {
	var myStringExample = string(`as`)
	println(myStringExample)

	http.HandleFunc("/", handleSlash)
	http.ListenAndServe(":8080", nil)
}

func handleSlash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	fmt.Println(r.URL.Path)

	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to home</h1>")

	} else if r.URL.Path == "/name" {
		fmt.Fprintf(w, "<h1>My name is Rohit</h1>")

	} else {
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(
			w,
			`<h1>404! couldn't find the page you requested</h1>
			<p>Please email us at
				<a href="mailto:hegde.rohit7@gmail.com">hegde.rohit7@gmail.com</a>
			</p>`)
	}
}
