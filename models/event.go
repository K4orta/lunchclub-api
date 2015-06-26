package models

import (
	"github.com/k4orta/lunchclub-api/storage/types"
	"time"
)

type EventId int

type Event struct {
	Id          EventId
	Title       string
	Slug        string
	StartTime   time.Time
	EndTime     time.Time
	OrangizerId UserId
	Rsvps       storage.IntList
}
