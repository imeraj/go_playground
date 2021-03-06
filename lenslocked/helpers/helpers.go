package helpers

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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

func ParseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}

func Remember(w http.ResponseWriter, user *models.User) {
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

func GetCookie(r *http.Request) (map[string]string, error) {
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

func imagePath(galleryID uint) string {
	return filepath.Join("public", "images", "galleries", fmt.Sprintf("%v", galleryID))
}

func CreateGalleryPath(galleryID uint) (string, error) {
	galleryPath := imagePath(galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}
