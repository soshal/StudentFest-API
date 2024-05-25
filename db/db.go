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

	createuser := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createuser)
		
	if err != nil {
		panic("Failed to create table")
	}
	
	createtable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATE ,
		user_id INTEGER ,
		FK(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createtable)

	if err != nil {

		panic("Failed to create table")
	}




}