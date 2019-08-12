package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/repositories"
)

type GalleryService struct {
	repo *repositories.GalleryRepo
}

func NewGalleryService() *GalleryService {
	repo := repositories.NewGalleryRepo()

	return &GalleryService{
		repo: repo,
	}
}

func (gs *GalleryService) Create(gallery *models.Gallery) error {
	return gs.repo.Create(gallery)
}
