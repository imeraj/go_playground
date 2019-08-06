package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *utils.Db) (*UserService, error) {
	return &UserService{db: db.Db}, nil
}

func (us *UserService) Create(user *models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}
