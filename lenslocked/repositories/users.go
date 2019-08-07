package repositories

import (
	"github.com/imeraj/go_playground/lenslocked/models"
	"github.com/imeraj/go_playground/lenslocked/utils/datastore"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	db := datastore.NewDB()
	return &UserRepo{db: db.Db}
}

func (repo *UserRepo) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepo) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepo) UserByEmail(email string) (*models.User, error) {
	var user models.User
	db := repo.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

func first(db *gorm.DB, user *models.User) error {
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		return models.ErrNotFound
	}
	return err
}
