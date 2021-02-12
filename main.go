package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	homeView     *View
	contactView  *View
	nameView     *View
	notFoundView *View
)

func main() {

	homeView = NewView("views/home.gohtml")
	contactView = NewView("views/contact.gohtml")
	nameView = NewView("views/name.gohtml")
	notFoundView = NewView("views/notfound.gohtml")

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

	_ = homeView.Template.Execute(w, nil)
}

func name(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)

	nameView.Template.Execute(w, nil)
}

func notFoundCustomPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	notFoundView.Template.Execute(w, nil)
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

	contactView.Template.Execute(w, nil)
}
