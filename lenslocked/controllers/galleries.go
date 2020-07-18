package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/imeraj/go_playground/lenslocked/helpers"

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
	EditView  *views.View
	gs        *services.GalleryService
}

type GalleryFrom struct {
	Title  string `schema:"title" validate:"required"`
	Errors map[string]string
}

type GalleryEditFrom struct {
	Gallery *models.Gallery
	Errors  map[string]string
}

const (
	maxMultiPartMemory = 1 << 20
)

func NewGallery() *Galleries {
	gs := services.NewGalleryService()

	return &Galleries{
		NewView:   views.NewView("bootstrap", "galleries/new"),
		ShowView:  views.NewView("bootstrap", "galleries/show"),
		IndexView: views.NewView("bootstrap", "galleries/index"),
		EditView:  views.NewView("bootstrap", "galleries/edit"),
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
	if err := helpers.ParseForm(r, &form); err != nil {
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
	gallery, err := g.galleryByID(w, r)
	if err != nil {
		return
	}

	g.ShowView.Render(w, gallery)
}

func (g *Galleries) Edit(w http.ResponseWriter, r *http.Request) {
	gallery, err := g.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "You do not have permission to edit this gallery.", http.StatusForbidden)
		return
	}

	var form GalleryEditFrom
	form.Gallery = gallery

	g.EditView.Render(w, form)
}

func (g *Galleries) Update(w http.ResponseWriter, r *http.Request) {
	gallery, err := g.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "You do not have permission to edit this gallery.", http.StatusForbidden)
		return
	}

	validationErrors := &ValidationErrors{}
	validationErrors.Errors = make(map[string]string)

	var form GalleryFrom
	if err := helpers.ParseForm(r, &form); err != nil {
		panic(err)
	}

	normalizeGalleryForm(&form)
	if validateForm(form, validationErrors) == false {
		var form1 GalleryEditFrom
		form1.Gallery = gallery
		form1.Gallery.Title = form.Title
		form1.Errors = validationErrors.Errors

		g.EditView.Render(w, form1)
		return
	}

	gallery.Title = form.Title
	err = g.gs.Update(gallery)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/galleries", http.StatusSeeOther)
}

func (g *Galleries) Delete(w http.ResponseWriter, r *http.Request) {
	gallery, err := g.galleryByID(w, r)
	if err != nil {
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

func (g *Galleries) ImageUpload(w http.ResponseWriter, r *http.Request) {
	gallery, err := g.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "You do not have upload permission for this gallery.", http.StatusForbidden)
		return
	}

	err = r.ParseMultipartForm(maxMultiPartMemory)
	if err != nil {
		g.EditView.Render(w, gallery)
		return
	}

	galleryPath, err := helpers.CreateGalleryPath(gallery.ID)
	if err != nil {
		g.EditView.Render(w, gallery)
		return
	}

	files := r.MultipartForm.File["images"]
	err = g.gs.ProcessImages(galleryPath, files)
	if err != nil {
		g.EditView.Render(w, gallery)
		return
	}

	url := fmt.Sprintf("/galleries/%v/edit", gallery.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (g *Galleries) ImageDelete(w http.ResponseWriter, r *http.Request) {
	gallery, err := g.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "You do not have edit permission for this gallery.", http.StatusForbidden)
		return
	}

	filename := mux.Vars(r)["filename"]
	i := models.Image{
		Filename:  filename,
		GalleryID: gallery.ID,
	}

	err = g.gs.DeleteImage(gallery, &i)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/galleries/%v/edit", gallery.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (g *Galleries) galleryByID(w http.ResponseWriter, r *http.Request) (*models.Gallery, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, models.ErrInvalidGalleryID.Error(), http.StatusNotFound)
		return nil, err
	}

	gallery, err := g.gs.ByID(uint(id))
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return nil, err
	}

	images, _ := g.gs.ByGalleryID(gallery.ID)
	gallery.Images = images
	return gallery, nil
}
