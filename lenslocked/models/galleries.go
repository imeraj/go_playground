package models

import (
	"time"
)

type Gallery struct {
	ID uint

	UserID uint
	Title  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GalleryRepo interface {
	Create(gallery *Gallery) error
}
