package models

import (
	"time"

	"github.com/k4orta/lunchclub-api/storage/types"
)

// EventID is a unique identifier
type EventID int

//Event is model which tracks the what, when and who
type Event struct {
	ID          EventID `json:"id" db:"id"`
	Title       string
	Slug        string
	StartTime   time.Time
	EndTime     time.Time
	OrangizerID UserId
	Rsvps       types.IntList
}
