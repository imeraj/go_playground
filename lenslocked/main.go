package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/controllers"
	middlewares "github.com/imeraj/go_playground/lenslocked/middlewares"
)

var errorC *controllers.Errors
var staticC *controllers.Static
var userC *controllers.Users
var sessionC *controllers.Sessions
var galleriesC *controllers.Galleries

var authMw *middlewares.Auth

func init() {
	errorC = controllers.NewErrors()
	staticC = controllers.NewStatic()
	userC = controllers.NewUser()
	sessionC = controllers.NewSession()
	galleriesC = controllers.NewGallery()

	authMw = middlewares.NewAuth(sessionC)
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.Handle("/login", sessionC.LoginView).Methods("GET")
	r.HandleFunc("/login", sessionC.Login).Methods("POST")

	r.HandleFunc("/galleries/new", authMw.ApplyFn(galleriesC.New)).Methods("GET")
	r.HandleFunc("/galleries", authMw.ApplyFn(galleriesC.Create)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(errorC.NotFound)

	http.ListenAndServe(":8080", r)
}
