package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/controllers"
)

var errorC *controllers.Errors
var staticC *controllers.Static
var userC *controllers.User

func init() {
	errorC = controllers.NewErrors()
	staticC = controllers.NewStatic()
	userC = controllers.NewUser()
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(errorC.NotFound)

	http.ListenAndServe(":8080", r)
}
