package controllers

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator"
	"github.com/imeraj/go_playground/lenslocked/models"
)

var default_lang = "en"

var validate *validator.Validate
var uni *ut.UniversalTranslator

var enTranslations = []struct {
	tag     string
	message string
}{
	{
		tag:     "alphanum",
		message: fmt.Sprintf("is not valid, must be alphanumeric"),
	},
	{
		tag:     "email",
		message: fmt.Sprintf("is not a valid email address"),
	},
	{
		tag:     "min",
		message: fmt.Sprintf("'{0}' is less than minimum length allowed '{1}'"),
	},
	{
		tag:     "max",
		message: fmt.Sprintf("'{0}' is more than maximum length allowed '{1}'"),
	},
	{
		tag:     "required",
		message: fmt.Sprintf("is a required field"),
	},
}

func init() {
	en := en_US.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()

	for _, t := range enTranslations {
		err := validate.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.message), translateFunc)
		if err != nil {
			panic(err)
		}
	}
}

func registrationFunc(tag string, message string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, message, true); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), reflect.ValueOf(fe.Value()).String(), fe.Param())
	if err != nil {
		return fe.(error).Error()
	}
	return t
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func validateForm(form interface{}, validationErrors *models.ValidationErrors) bool {
	errs := validate.Struct(form)
	if errs == nil {
		return true
	}

	trans, _ := uni.GetTranslator(default_lang)

	for _, err := range errs.(validator.ValidationErrors) {
		jsonKey := err.Field()
		fieldName, _ := trans.T(jsonKey)
		message := strings.Replace(err.Translate(trans), jsonKey, fieldName, -1)
		jsonKey = jsonKey[0:len(jsonKey)]
		validationErrors.Errors[jsonKey] = message
	}

	return len(validationErrors.Errors) == 0
}

func normalizeSignUpForm(form *SignupForm) error {
	form.Name = strings.TrimSpace(form.Name)

	form.Email = strings.ToLower(form.Email)
	form.Email = strings.TrimSpace(form.Email)

	return nil
}
