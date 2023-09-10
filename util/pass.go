package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed hash %w", err)
	}

	return string(h), nil
}

func CheckPass(pass string, h string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(pass))
}
