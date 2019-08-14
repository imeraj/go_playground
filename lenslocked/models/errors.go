package models

import "errors"

var (
	ErrNotFound         = errors.New("models: User not found")
	ErrInvalidPassword  = errors.New("models: Incorrect password provided")
	ErrInvalidGalleryID = errors.New("models: Incorrect gallery ID provided")
)
