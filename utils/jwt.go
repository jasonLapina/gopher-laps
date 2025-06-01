package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
)

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
	})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwtKey := os.Getenv("JWT_KEY")

	return token.SignedString(jwtKey)

}
