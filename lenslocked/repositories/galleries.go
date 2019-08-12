package repositories

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/utils/datastore"

	"github.com/jinzhu/gorm"
)

type GalleryRepo struct {
	db *gorm.DB
}

func NewGalleryRepo() *GalleryRepo {
	db := datastore.NewDB()
	return &GalleryRepo{db: db.Db}
}

func (repo *GalleryRepo) Create(gallery *models.Gallery) error {
	return repo.db.Create(gallery).Error
}
