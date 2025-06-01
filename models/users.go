package models

import (
	"awesomeProject/database"
	"awesomeProject/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPw := utils.HashPassword(u.Password)

	res, err := statement.Exec(u.Email, hashedPw)

	if err != nil {
		return err
	}

	userId, _ := res.LastInsertId()

	u.ID = userId

	return nil

}

func (u User) ValidateCreds() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := database.DB.QueryRow(query, u.Email)

	var retrievedPW string

	err := row.Scan(u.ID, &retrievedPW)

	if err != nil {
		return err
	}

	err = utils.ComparePassword(retrievedPW, u.Password)

	if err != nil {
		return err
	}

	return nil
}
