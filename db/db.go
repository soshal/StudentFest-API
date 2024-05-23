package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB()  {
	DB, err := sql.Open("sqlite3", "./college.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	DB.SetConnMaxIdleTime(5)
	DB.SetMaxOpenConns(10)

}

func createtable() {

	createtable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATE ,
		user_id INTEGER 
	)`

	_, err := DB.Exec(createtable)

	if err != nil {

		panic("Failed to create table")
	}




}