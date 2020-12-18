package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

func main() {
	var err error

	homeTemplate, err = template.ParseFiles("views/home.gohtml")

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFoundCustomPage)

	router.HandleFunc("/", home)
	router.HandleFunc("/name", name)
	router.HandleFunc("/faq", faqPage)

	//http.HandleFunc("/", handleSlash)
	_ = http.ListenAndServe(":8081", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	_ = homeTemplate.Execute(w, nil)
	//_, _ = fmt.Fprintf(w, "<h1>Welcome to home</h1>")
}

func name(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprintf(w, "<h1>My name is Rohit</h1>")
}

func notFoundCustomPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	_, _ = fmt.Fprintf(
		w,
		`<h1>404! couldn't find the page you requested</h1>
			<p>Please email me at
				<a href="mailto:hegde.rohit7@gmail.com">
					hegde.rohit7@gmail.com
				</a>
			</p>`)
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	fmt.Fprintf(w,
		`<h1>FAQ (Frequently asked questions)</h1>
			<p>Kuch puchna hai toh idhar mail karne ka - 
				<a href="mailto:hegde.rohit7@gmail.com">
					hegde.rohit7@gmail.com
				</a>
				Kya samjha!?
			</p>`)
}
