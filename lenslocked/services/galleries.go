package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	galleriesrepo "github.com/imeraj/go_playground/lenslocked/repositories/gallery"
)

type GalleryService struct {
	repo *galleriesrepo.GalleryRepo
}

func NewGalleryService() *GalleryService {
	repo := galleriesrepo.NewGalleryRepo()

	return &GalleryService{
		repo: repo,
	}
}

func (gs *GalleryService) Create(gallery *models.Gallery) error {
	return gs.repo.Create(gallery)
}

func (gs *GalleryService) ByUserID(userID uint) ([]models.Gallery, error) {
	return gs.repo.ByUserID(userID)
}

func (gs *GalleryService) ByID(id uint) (*models.Gallery, error) {
	return gs.repo.ByID(id)
}
