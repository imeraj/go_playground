package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID uint

	Name         string
	Email        string
	Password     string `gorm:"-"`
	PasswordHash string

	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	ErrNotFound        = errors.New("models: User not found")
	ErrInvalidPassword = errors.New("models: Incorrect password provided")
)

func (u *User) First(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
