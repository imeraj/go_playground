package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	db *gorm.DB
}

func NewSessionService(db *utils.Db) (*SessionService, error) {
	return &SessionService{db: db.Db}, nil
}

func (ss *SessionService) Authenticate(email, password string) (*models.User, error) {
	foundUser, err := ss.UserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, models.ErrInvalidPassword
	default:
		return nil, err
	}
}

func (ss *SessionService) UserByEmail(email string) (*models.User, error) {
	var user models.User
	db := ss.db.Where("email = ?", email)
	err := user.First(db, &user)
	return &user, err
}
