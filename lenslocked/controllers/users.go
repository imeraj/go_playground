package controllers

import (
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/helpers"

	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/services"
	"github.com/imeraj/go_playground/lenslocked/views"
)

type Users struct {
	NewView *views.View
	us      *services.UserService
}

func NewUsers() *Users {
	us := services.NewUserService()

	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, r, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	validationErrors := &helpers.ValidationErrors{}
	validationErrors.Errors = make(map[string]string)

	var form helpers.SignupForm
	var vd views.Data

	if err := helpers.ParseForm(r, &form); err != nil {
		panic(err)
	}

	helpers.NormalizeSignUpForm(&form)
	if helpers.ValidateForm(form, validationErrors) == false {
		form.Errors = validationErrors.Errors
		u.NewView.Render(w, r, form)
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		vd.Yield = form
		vd.SetAlert(err.Error(), views.AlertLvlError)
		u.NewView.Render(w, r, vd)
		return
	}

	vd.Yield = form
	vd.SetAlert("User created successfully.", views.AlertLvlSuccess)
	u.NewView.Render(w, r, vd)
}
