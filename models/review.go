package models

import "time"

// ReviewID is a unique identifier
type ReviewID int

// Review is a rating given to a location by a user
type Review struct {
	ID       ReviewID   `json:"id" db:"id"`
	Location LocationID `json:"location" db:"location_id"`
	Rating   float32    `json:"rating" db:"rating"`
	Reviewer UserID     `json:"reviewer" db:"reviewer"`
	Remarks  string     `json:"remarks" db:"remarks"`
	Added    time.Time  `json:"added" db:"added"`
	Edited   time.Time  `json:"edited" db:"edited"`
}
