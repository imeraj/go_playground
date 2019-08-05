package services

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/utils"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *utils.Db) (*UserService, error) {
	return &UserService{db: db.Db}, nil
}

func (us *UserService) Create(user *models.User) error {
	return us.db.Create(user).Error
}
