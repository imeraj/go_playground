package models

import (
	"time"
	"errors"
)

var (
	ErrNotFound         = errors.New("models: User not found")
	ErrInvalidPassword  = errors.New("models: Incorrect password provided")
)

type User struct {
	ID uint

	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string
	Remember     string `gorm:"-"`
	RememberHash string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	Create(user *User) error
	Update(user *User) error
	ByEmail(email string) (*User, error)
	ByRemember(remember string) (*User, error)
}
