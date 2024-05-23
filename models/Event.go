package main 

import "time"


type Event struct {
	Id int
	name string
	description string
	Date time.Time
	Location string
	UserId int
}
var event = []Event{ }

func (e *Event) Save() {


	event = append(event, *e)
}


