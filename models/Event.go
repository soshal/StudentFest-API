package models

import (
	"college/db"
	"time"

	"golang.org/x/tools/go/analysis/passes/defers"
)

type Event struct {
	Id          int64
	Name        string  `Binding:"required"`
	Description string	`Binding:"required"`

	Date        time.Time
	Location    string `Binding:"required"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() {
	query := `INSERT INTO events(name, description, location, date, user_id) VALUES(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserId)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	e.Id = id
	return err
}

func GetEvents() []Event {
	return events
}

