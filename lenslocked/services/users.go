package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	usersrepo "github.com/imeraj/go_playground/lenslocked/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

const hmacSecretKey = "secret-hmac-key" // ideally should not be stored in code

type UserService struct {
	repo *usersrepo.UserRepo
}

func NewUserService() *UserService {
	repo := usersrepo.NewUserRepo()

	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Create(user *models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	return us.repo.Create(user)
}
