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
	userC = controllers.NewUsers()
	sessionC = controllers.NewSession()
	galleriesC = controllers.NewGallery()

	authMw = middlewares.NewAuth(sessionC)
}

func main() {
	r := mux.NewRouter()

	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	imageHandler := http.FileServer(http.Dir("public/images/"))
	r.PathPrefix("/public/images/").Handler(http.StripPrefix("/public/images/", imageHandler))

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.Handle("/login", sessionC.LoginView).Methods("GET")
	r.HandleFunc("/login", sessionC.Login).Methods("POST")

	r.HandleFunc("/galleries/new", authMw.ApplyFn(galleriesC.New)).Methods("GET")
	r.HandleFunc("/galleries", authMw.ApplyFn(galleriesC.Create)).Methods("POST")
	r.HandleFunc("/galleries", authMw.ApplyFn(galleriesC.Index)).Methods("GET")
	r.HandleFunc("/galleries/{id:[0-9]+}", authMw.ApplyFn(galleriesC.Show)).Methods("GET")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete", authMw.ApplyFn(galleriesC.Delete)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", authMw.ApplyFn(galleriesC.Edit)).Methods("GET")
	r.HandleFunc("/galleries/{id:[0-9]+}/update", authMw.ApplyFn(galleriesC.Update)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/images", authMw.ApplyFn(galleriesC.ImageUpload)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/images/{filename}/delete", authMw.ApplyFn(galleriesC.ImageDelete)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(errorC.NotFound)
	http.ListenAndServe(":8080", r)
}
