package models

import (
	"encoding/json"
	"errors"
	"log"
)

const (
	oneshot = 0
	monthly = 24 * 60 * 60
)

// Event - Model of a basic event
type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Freq  int32  `json:"freq"`
	Desc  string `json:"desc"`
}

// GetDefaultEvents .. sample events to begin with
func GetDefaultEvents() []Event {
	events := []Event{
		{
			ID:    1,
			Title: "Pay Electricity Bill",
			Freq:  oneshot,
			Desc:  "One time bill",
		},
		{
			ID:    2,
			Title: "Pay Gas Bill",
			Freq:  monthly,
			Desc:  "One time bill",
		},
	}
	return events
}

// GetEventsByID .. return event for specific ID
func GetEventsByID(eid int) (*Event, error) {
	events := GetDefaultEvents()

	for _, e := range events {
		if e.ID == eid {
			return &e, nil
		}
	}

	return nil, errors.New("Event not found")
}

// Encode .. returns json encoded event
func (e Event) Encode() []byte {
	bytes, err := json.Marshal(&e)
	if err != nil {
		log.Fatalln(err)
	}

	return bytes
}
