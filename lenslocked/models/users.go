package models

import (
	"errors"
	"time"
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

type UserRepo interface {
	Create(*User) error
	UserByEmail(email string) (*User, error)
}

var (
	ErrNotFound        = errors.New("models: User not found")
	ErrInvalidPassword = errors.New("models: Incorrect password provided")
)
