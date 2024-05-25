package models

import (
	"college/db"
	"time"

	
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
	if err != nil {
		panic(err)
	}
}

func GetEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		e := Event{}
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.Date, &e.UserId)
		if err != nil {
			 return nil, err
		}
		events = append(events, e)
	}
	return events,nil
}

func GetEventbyId(id int64) (Event, error) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	e := Event{}
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.Date, &e.UserId)
	if err != nil {
		return e, err
	}
	return e, nil
}

func (e Event) UpdateEvent() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.Id)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(e.Id)
	if err != nil {
		return err
	}
	return nil
}	



