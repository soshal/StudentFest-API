package models

import (
	"college/db"
	"college/utils"
	"time"
)


type User struct {
	
	Id          int64
	Email       string `Binding:"required"`
	Password    string `Binding:"required"`
	
}	


func (u *User) Save() {
	query := `INSERT INTO users(email, password) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	
	
	password, err :=utils.HashPassword(u.Password) // hashing the passowrd

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(u.Email, password)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	u.Id = id
	if err != nil {
		panic(err)
	}
}


