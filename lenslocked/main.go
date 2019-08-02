package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/views"
)

var homeView *views.View
var contactView *views.View
var notFoundView *views.View

func init() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	notFoundView = views.NewView("bootstrap", "views/errors/404.gohtml")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	must(notFoundView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8080", r)
}
