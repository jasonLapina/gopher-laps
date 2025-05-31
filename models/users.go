package models

import "awesomeProject/database"

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
	res, err := statement.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, _ := res.LastInsertId()

	u.ID = userId

	return nil

}
