package controllers

import (
	"fmt"
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/views"
)

type User struct {
	NewView *views.View
}

func NewUser() *User {
	return &User{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml")}
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response")
}
