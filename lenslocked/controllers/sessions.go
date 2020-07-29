package controllers

import (
	"net/http"
	"time"

	"github.com/imeraj/go_playground/lenslocked/utils/rand"

	"github.com/imeraj/go_playground/lenslocked/context"
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

	var vd views.Data
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
		vd.SetAlert("Invalid email address", views.AlertLvlError)
	case models.ErrInvalidPassword:
		vd.SetAlert("Invalid password provided", views.AlertLvlError)
	default:
		vd.SetAlert(err.Error(), views.AlertLvlError)
	}

	vd.Yield = form
	s.LoginView.Render(w, r, vd)
}

func (s *Sessions) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	user := context.User(r.Context())
	token, _ := rand.RememberToken()
	user.Remember = token
	s.ss.Update(user)

	http.Redirect(w, r, "/", http.StatusFound)
}
