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

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, keyFunc)

	if err != nil {
		return 0, err
	}

	isValid := parsedToken.Valid

	if !isValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil

}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, errors.New("Invalid token")
	}

	return os.Getenv("JWT_KEY"), nil
}
