package helpers

import "github.com/imeraj/go_playground/lenslocked/models"

type SignupForm struct {
	Name     string `schema:"name" validate:"alphanum,required"`
	Email    string `schema:"email" validate:"email,required"`
	Password string `schema:"password" validate:"min=3,max=8,required"`
	Errors   map[string]string
}

type GalleryEditFrom struct {
	Gallery *models.Gallery
	Errors  map[string]string
}

type GalleryFrom struct {
	Title  string `schema:"title" validate:"required"`
	Errors map[string]string
}
