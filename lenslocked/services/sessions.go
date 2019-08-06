package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/repositories"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	repo *repositories.UserRepo
}

func NewSessionService() *SessionService {
	repo := repositories.NewUserRepo()
	return &SessionService{repo: repo}
}

func (ss *SessionService) Authenticate(email, password string) (*models.User, error) {
	foundUser, err := ss.userByEmail(email)
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

func (ss *SessionService) userByEmail(email string) (*models.User, error) {
	return ss.repo.UserByEmail(email)
}
