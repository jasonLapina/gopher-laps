package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwtKey := os.Getenv("JWT_KEY")

	return token.SignedString([]byte(jwtKey))

}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, keyFunc)

	if err != nil {
		return err
	}

	isValid := parsedToken.Valid

	if !isValid {
		return errors.New("Invalid token")
	}

	return nil

}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, errors.New("Invalid token")
	}

	return os.Getenv("JWT_KEY"), nil
}
