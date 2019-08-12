package controllers

import (
	"fmt"
	"net/http"

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

type LoginForm struct {
	Email    string `schema:"email" validate:"email,required"`
	Password string `schema:"password" validate:"required"`
	Errors   map[string]string
}

func NewSession() *Sessions {
	ss := services.NewSessionService()

	return &Sessions{
		LoginView: views.NewView("bootstrap", "sessions/login"),
		ss:        ss,
	}
}

func (s *Sessions) GetCookie(r *http.Request) (map[string]string, error) {
	return getCookie(r)
}

func (s *Sessions) GetSessionService() *services.SessionService {
	return s.ss
}

func (s *Sessions) Login(w http.ResponseWriter, r *http.Request) {
	validationErrors := &ValidationErrors{}
	validationErrors.Errors = make(map[string]string)
	form := LoginForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	if validateForm(form, validationErrors) == false {
		form.Errors = validationErrors.Errors
		s.LoginView.Render(w, form)
		return
	}

	user, err := s.ss.Authenticate(form.Email, form.Password)
	switch err {
	case nil:
		remember(w, user)
		http.Redirect(w, r, "/galleries/new", http.StatusSeeOther)
	case models.ErrNotFound:
		fmt.Fprintf(w, "Invalid email address.")
	case models.ErrInvalidPassword:
		fmt.Fprintf(w, "Invalid password provided.")
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
