package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) string {

	b, _ := bcrypt.GenerateFromPassword([]byte(pw), 10)

	return string(b)
}

func ComparePassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
