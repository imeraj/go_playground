package galleriesrepo

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

func (repo *GalleryRepo) ByUserID(userID uint) ([]models.Gallery, error) {
	var galleries []models.Gallery
	db := repo.db.Where("user_id = ?", userID)
	if err := db.Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}

func (repo *GalleryRepo) ByID(galleryID uint) (*models.Gallery, error) {
	var gallery models.Gallery
	db := repo.db.Where("id = ?", galleryID)
	err := first(db, &gallery)
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func first(db *gorm.DB, gallery *models.Gallery) error {
	err := db.First(gallery).Error
	if err == gorm.ErrRecordNotFound {
		return models.ErrNotFound
	}
	return err
}
