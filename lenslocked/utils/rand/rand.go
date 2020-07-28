package rand

import (
	"encoding/base64"
	"math/rand"
)

const RememberTokenBytes = 32

func RememberToken() (string, error) {
	return randString(RememberTokenBytes)
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func randString(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
