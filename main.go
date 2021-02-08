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
	nameTemplate *template.Template
	notFoundTemplate *template.Template
)

func main() {
	var err error
	var contactTemplateErr error

	homeTemplate, err = template.ParseFiles("views/home.gohtml", "views/layout/footer.html")
	contactTemplate, contactTemplateErr = template.ParseFiles("views/contact.gohtml", "views/layout/footer.html")
	nameTemplate, _ = template.ParseFiles("views/name.gohtml", "views/layout/footer.html")
	notFoundTemplate, _ = template.ParseFiles("views/notfound.gohtml", "views/layout/footer.html")

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
	w.WriteHeader(http.StatusOK)

	nameTemplate.Execute(w, nil)
}

func notFoundCustomPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	notFoundTemplate.Execute(w, nil)
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
