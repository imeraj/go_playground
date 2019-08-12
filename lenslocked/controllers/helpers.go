package controllers

import (
	"encoding/hex"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
	"github.com/imeraj/go_playground/lenslocked/models"
)

// ideally should not be stored in source code
var hashKey = "d93205c96caf4258a361eb5a0209075c95e1c2a854f9c90b9f73f3a1a2da058e"
var blockKey = "a50ccfe95fe55a2cd592270d4a4ff3b9a41fa8725ed15368394141154097d887"

var s *securecookie.SecureCookie

func init() {
	hKey, _ := hex.DecodeString(hashKey)
	bKey, _ := hex.DecodeString(blockKey)

	s = securecookie.New(hKey, bKey)
}

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}

func remember(w http.ResponseWriter, user *models.User) {
	setCookie(w, user)
}

func setCookie(w http.ResponseWriter, user *models.User) {
	value := map[string]string{
		"remember_token": user.Remember,
	}

	if encoded, err := s.Encode("lenslocked", value); err == nil {
		cookie := &http.Cookie{
			Name:  "lenslocked",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getCookie(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("lenslocked")
	if err != nil {
		return nil, err
	}

	value := make(map[string]string)
	if err = s.Decode("lenslocked", cookie.Value, &value); err == nil {
		return value, nil
	}
	return nil, err
}
