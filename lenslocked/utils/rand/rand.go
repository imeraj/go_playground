package rand

import (
	"encoding/base64"
	"math/rand"
)

const RememberTokenBytes = 32

func RememberToken() (string, error) {
	return randString(RememberTokenBytes)
}

func randBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func randString(nBytes int) (string, error) {
	b, err := randBytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
