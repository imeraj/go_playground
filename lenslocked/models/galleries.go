package models

import (
	"time"
)

type Gallery struct {
	ID uint

	UserID uint
	Title  string
	Images []string `gorm:"-"`

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

func (g *Gallery) ImagesSplitN(n int) [][]string {
	// Create out 2D slice
	ret := make([][]string, n)
	// Create the inner slices - we need N of them, and we will
	// start them with a size of 0.
	for i := 0; i < n; i++ {
		ret[i] = make([]string, 0)
	}
	// Iterate over our images, using the index % n to determine
	// which of the slices in ret to add the image to.
	for i, img := range g.Images {
		// % is the remainder operator in Go
		// eg:
		//    0%3 = 0
		//    1%3 = 1
		//    2%3 = 2
		//    3%3 = 0
		//    4%3 = 1
		//    5%3 = 2
		bucket := i % n
		ret[bucket] = append(ret[bucket], img)
	}
	return ret
}
