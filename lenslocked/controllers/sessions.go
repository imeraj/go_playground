package controllers

import (
	"fmt"
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/helpers"

	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/services"
	"github.com/imeraj/go_playground/lenslocked/utils/hash"
	"github.com/imeraj/go_playground/lenslocked/views"
)

type Sessions struct {
	LoginView *views.View
	ss        *services.SessionService
	hmac      hash.HMAC
}

func NewSession() *Sessions {
	ss := services.NewSessionService()

	return &Sessions{
		LoginView: views.NewView("bootstrap", "sessions/login"),
		ss:        ss,
	}
}

func (s *Sessions) GetCookie(r *http.Request) (map[string]string, error) {
	return helpers.GetCookie(r)
}

func (s *Sessions) GetSessionService() *services.SessionService {
	return s.ss
}

func (s *Sessions) Login(w http.ResponseWriter, r *http.Request) {
	validationErrors := &helpers.ValidationErrors{}
	validationErrors.Errors = make(map[string]string)
	form := helpers.LoginForm{}

	if err := helpers.ParseForm(r, &form); err != nil {
		panic(err)
	}

	if helpers.ValidateForm(form, validationErrors) == false {
		form.Errors = validationErrors.Errors
		s.LoginView.Render(w, r, form)
		return
	}

	user, err := s.ss.Authenticate(form.Email, form.Password)
	switch err {
	case nil:
		helpers.Remember(w, user)
		http.Redirect(w, r, "/galleries", http.StatusSeeOther)
	case models.ErrNotFound:
		fmt.Fprintf(w, "Invalid email address.")
	case models.ErrInvalidPassword:
		fmt.Fprintf(w, "Invalid password provided.")
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
