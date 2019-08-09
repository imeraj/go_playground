package controllers

import (
	"strings"

	validator "github.com/go-playground/validator"
	"github.com/imeraj/go_playground/lenslocked/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateForm(form interface{}) error {
	return validate.Struct(form)
}

func normalizeEmail(user *models.User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}
