package utils

import "math/rand"

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
