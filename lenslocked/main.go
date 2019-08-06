package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/controllers"
)

var errorC *controllers.Errors
var staticC *controllers.Static
var userC *controllers.Users
var sessionC *controllers.Sessions

func init() {
	errorC = controllers.NewErrors()
	staticC = controllers.NewStatic()
	userC = controllers.NewUser()
	sessionC = controllers.NewSession()
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.Handle("/login", sessionC.LoginView).Methods("GET")
	r.HandleFunc("/login", sessionC.Login).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(errorC.NotFound)

	http.ListenAndServe(":8080", r)
}
