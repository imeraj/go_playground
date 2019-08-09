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
}

func NewSession() *Sessions {
	ss := services.NewSessionService()

	return &Sessions{
		LoginView: views.NewView("bootstrap", "sessions/login"),
		ss:        ss,
	}
}

func (s *Sessions) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	if err := validateForm(form); err != "" {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}

	user, err := s.ss.Authenticate(form.Email, form.Password)
	switch err {
	case nil:
		remember(w, user)
		fmt.Fprintf(w, "Login successful.")
	case models.ErrNotFound:
		fmt.Fprintf(w, "Invalid email address.")
	case models.ErrInvalidPassword:
		fmt.Fprintf(w, "Invalid password provided.")
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
