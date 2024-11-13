package service

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortenerLink(link string) (string, error) {

	bytes := make([]byte, len(link))
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	shortenerLink := "http://generated-link/" + base64.URLEncoding.EncodeToString(bytes)[:len(link)]

	return shortenerLink, nil
}
