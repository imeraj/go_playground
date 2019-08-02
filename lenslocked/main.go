package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/controllers"
	"github.com/imeraj/go_playground/lenslocked/views"
)

var homeView *views.View
var contactView *views.View
var notFoundView *views.View

var userC *controllers.User

func init() {
	userC = controllers.NewUser()

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

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")

	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":8080", r)
}
