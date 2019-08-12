package controllers

import (
	"fmt"
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/services"
	"github.com/imeraj/go_playground/lenslocked/views"
)

type Users struct {
	NewView *views.View
	us      *services.UserService
}

type SignupForm struct {
	Name     string `schema:"name" validate:"alphanum,required"`
	Email    string `schema:"email" validate:"email,required"`
	Password string `schema:"password" validate:"min=3,max=8,required"`
	Errors   map[string]string
}

func NewUser() *Users {
	us := services.NewUserService()

	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	validationErrors := &models.ValidationErrors{}
	validationErrors.Errors = make(map[string]string)

	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	normalizeSignUpForm(&form)
	if validateForm(form, validationErrors) == false {
		form.Errors = validationErrors.Errors
		u.NewView.Render(w, form)
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created!")
}
