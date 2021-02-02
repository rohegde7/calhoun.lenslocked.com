package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func main() {
	var err error
	var contactTemplateErr error

	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	contactTemplate, contactTemplateErr = template.ParseFiles("views/contact.gohtml")

	if err != nil && contactTemplateErr != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFoundCustomPage)

	router.HandleFunc("/", home)
	router.HandleFunc("/name", name)
	router.HandleFunc("/faq", faqPage)
	router.HandleFunc("/contact", contact)

	_ = http.ListenAndServe(":8080", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	_ = homeTemplate.Execute(w, nil)
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

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusAccepted)

	contactTemplate.Execute(w, nil)
}
