package models

import (
	"errors"
	"time"
)

var (
	ErrInvalidGalleryID = errors.New("models: Incorrect gallery ID provided")
)

type Gallery struct {
	ID uint

	UserID uint `gorm:"not null;index"`
	Title  string
	Images []Image `gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GalleryRepo interface {
	Create(gallery *Gallery) error
	ByUserID(userID uint) ([]Gallery, error)
	ByID(galleryID uint) (*Gallery, error)
	Delete(galleryID uint) error
	Update(gallery *Gallery) error
}

func (g *Gallery) ImagesSplitN(n int) [][]Image {
	ret := make([][]Image, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]Image, 0)
	}
	for i, img := range g.Images {
		bucket := i % n
		ret[bucket] = append(ret[bucket], img)
	}
	return ret
}
