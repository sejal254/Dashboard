package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt) + password

	return hashedPassword, nil
}

func VerifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}
