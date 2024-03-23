package models

import "time"

type Event struct {
	ID          int       `binding:"required"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

// Atteched to Event
func (e Event) Save() {
	// Save the event
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
