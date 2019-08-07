package services

import (
	"fmt"

	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/repositories"
	"github.com/imeraj/go_playground/lenslocked/utils/hash"
	"github.com/imeraj/go_playground/lenslocked/utils/rand"
	"golang.org/x/crypto/bcrypt"
)

const hmacSecretKey = "secret-hmac-key" // ideally should not be stored in code

type UserService struct {
	repo *repositories.UserRepo
	hmac hash.HMAC
}

func NewUserService() *UserService {
	hmac := hash.NewHMAC(hmacSecretKey)
	repo := repositories.NewUserRepo()

	return &UserService{
		repo: repo,
		hmac: hmac,
	}
}

func (us *UserService) Create(user *models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		user.RememberHash = us.hmac.Hash(user.Remember)
		fmt.Printf("%d", len(user.RememberHash))
	}

	return us.repo.Create(user)
}
