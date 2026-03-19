package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"`
}

var events = []Event{}

func (e Event) Save() Event {
	event := Event{
		ID:          len(events) + 1,
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		DateTime:    time.Now(),
		UserId:      1,
	}
	events = append(events, event)
	return event
}

func GetAllEvents() []Event {
	return events
}
