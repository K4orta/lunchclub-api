package models

import (
	"time"

	"github.com/k4orta/lunchclub-api/storage/types"
)

type ClubID int

// Club is the model and glues together users and events
type Club struct {
	ID               ClubID        `json:"id" db:"id"`
	Members          types.IntList `json:"members" db:"members"`
	Events           types.IntList `json:"events" db:"events"`
	DefaultEventDay  int           `json:"defaultEventDay" db:"default_event_day"`
	DefaultEventTime time.Time     `json:"defaultEventTime" db:"default_event_time"`
}
