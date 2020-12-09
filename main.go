package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/name", name)

	//http.HandleFunc("/", handleSlash)
	_ = http.ListenAndServe(":8081", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprintf(w, "<h1>Welcome to home</h1>")
}

func name(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprintf(w, "<h1>My name is Rohit</h1>")
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	_, _ = fmt.Fprintf(
		w,
		`<h1>404! couldn't find the page you requested</h1>
			<p>Please email us at
				<a href="mailto:hegde.rohit7@gmail.com">
					hegde.rohit7@gmail.com
				</a>
			</p>`)
}
