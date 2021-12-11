package utils

import (
	"fmt"
	"math/rand"
	"net/url"
)

const (
	asciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	asciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
)

func GenerateSlug() string {
	chars := []rune(asciiLowercase + asciiUppercase + digits + "_")

	s := make([]rune, 10)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}

func ValidateURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("'%s' is a invalid url", u)
	}

	return nil

}
