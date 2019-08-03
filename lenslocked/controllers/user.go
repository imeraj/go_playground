package controllers

import (
	"fmt"
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/views"
)

type User struct {
	NewView *views.View
}

type SingupForm struct {
	Email    string `schema: "email"`
	Password string `schema: "password"`
}

func NewUser() *User {
	return &User{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var form SingupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "email is: ", form.Email)
	fmt.Fprintln(w, "password is: ", form.Password)
}
