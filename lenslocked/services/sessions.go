package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	usersrepo "github.com/imeraj/go_playground/lenslocked/repositories/user"
	"github.com/imeraj/go_playground/lenslocked/utils/hash"
	"github.com/imeraj/go_playground/lenslocked/utils/rand"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	repo *usersrepo.UserRepo
	hmac hash.HMAC
}

func NewSessionService() *SessionService {
	hmac := hash.NewHMAC(hmacSecretKey)
	repo := usersrepo.NewUserRepo()

	return &SessionService{
		repo: repo,
		hmac: hmac,
	}
}

func (ss *SessionService) Authenticate(email, password string) (*models.User, error) {
	foundUser, err := ss.userByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
	switch err {
	case nil:
		return ss.remember(foundUser)
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, models.ErrInvalidPassword
	default:
		return nil, err
	}
}

func (ss *SessionService) userByEmail(email string) (*models.User, error) {
	return ss.repo.ByEmail(email)
}

func (ss *SessionService) remember(user *models.User) (*models.User, error) {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return nil, err
		}
		user.Remember = token
		user.RememberHash = ss.hmac.Hash(user.Remember)
		err = ss.repo.Update(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (ss *SessionService) ByRemember(token string) (*models.User, error) {
	user, err := ss.repo.ByRemember(ss.hmac.Hash(token))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ss *SessionService) Update(user *models.User) {
	ss.repo.Update(user)
}
