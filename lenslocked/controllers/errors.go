package controllers

import (
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/views"
)

type Errors struct {
	NotFoundView *views.View
}

func NewErrors() *Errors {
	return &Errors{
		NotFoundView: views.NewView("bootstrap", "errors/404"),
	}
}

func (e *Errors) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	must(e.NotFoundView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
