package models

import (
	"time"

	"github.com/k4orta/lunchclub-api/storage/types"
)

// EventID is a unique identifier
type EventID int

//Event is model which tracks the what, when and who
type Event struct {
	ID          EventID       `json:"id" db:"id"`
	Title       string        `json:"title" db:"title"`
	Slug        string        `json:"slug" db:"slug"`
	StartTime   time.Time     `json:"startTime" db:"start_time"`
	EndTime     time.Time     `json:"endTime" db:"end_time"`
	OrangizerID UserID        `json:"organizerID" db:"organizer_id"`
	LocationID  LocationID    `json:"locationID" db:"location_id"`
	RSVPs       types.IntList `json:"rsvps" db:"rsvps"`
}
