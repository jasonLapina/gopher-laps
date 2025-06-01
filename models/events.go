package models

import (
	"awesomeProject/database"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {

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

	e.ID = id
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

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (e Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
		`
	statement, err := database.DB.Prepare(query)
	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err

}

func (e Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	statement, err := database.DB.Prepare(query)
	defer statement.Close()

	if err != nil {
		return err
	}
	_, err = statement.Exec(e.ID)
	return err

}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)
	`
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)

	return err
}
