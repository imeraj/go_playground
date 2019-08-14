package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imeraj/go_playground/lenslocked/context"
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/services"
	"github.com/imeraj/go_playground/lenslocked/views"
)

type Galleries struct {
	NewView   *views.View
	ShowView  *views.View
	IndexView *views.View
	gs        *services.GalleryService
}

type GalleryFrom struct {
	Title  string `schema:"title" validate:"required"`
	Errors map[string]string
}

func NewGallery() *Galleries {
	gs := services.NewGalleryService()

	return &Galleries{
		NewView:   views.NewView("bootstrap", "galleries/new"),
		ShowView:  views.NewView("bootstrap", "galleries/show"),
		IndexView: views.NewView("bootstrap", "galleries/index"),
		gs:        gs,
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

	http.Redirect(w, r, "/galleries", http.StatusSeeOther)
}

func (g *Galleries) Index(w http.ResponseWriter, r *http.Request) {
	user := context.User(r.Context())
	galleries, err := g.gs.ByUserID(user.ID)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	g.IndexView.Render(w, galleries)
}

func (g *Galleries) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, models.ErrInvalidGalleryID.Error(), http.StatusNotFound)
		return
	}

	gallery, err := g.gs.ByID(uint(id))
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	g.ShowView.Render(w, gallery)
}

func (g *Galleries) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, models.ErrInvalidGalleryID.Error(), http.StatusNotFound)
		return
	}

	gallery, err := g.gs.ByID(uint(id))
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "You do not have permission to delete this gallery.", http.StatusForbidden)
		return
	}

	err = g.gs.Delete(gallery.ID)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/galleries", http.StatusSeeOther)
}
