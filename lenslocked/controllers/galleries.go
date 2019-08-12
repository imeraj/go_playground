package controllers

import (
	"fmt"
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/context"
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/services"
	"github.com/imeraj/go_playground/lenslocked/views"
)

type Galleries struct {
	NewView *views.View
	gs      *services.GalleryService
}

type GalleryFrom struct {
	Title  string `schema:"title" validate:"required"`
	Errors map[string]string
}

func NewGallery() *Galleries {
	gs := services.NewGalleryService()

	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
		gs:      gs,
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	validationErrors := &ValidationErrors{}
	validationErrors.Errors = make(map[string]string)

	var form GalleryFrom

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	normalizeGalleryForm(&form)
	if validateForm(form, validationErrors) == false {
		form.Errors = validationErrors.Errors
		g.NewView.Render(w, form)
		return
	}

	user := context.User(r.Context())
	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}

	if err := g.gs.Create(&gallery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Gallery created!")
}
