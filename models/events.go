package models

import (
	"awesomeProject/database"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {

	query := "INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)"

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	res, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	defer statement.Close()

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()

	e.ID = int(id)
	return nil

}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
