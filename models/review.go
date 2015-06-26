package models

import "time"

type ReviewID int

type Review struct {
	ID       ReviewID  `json:"id" db:"id"`
	Rating   float32   `json:"rating" db:"rating"`
	Reviewer UserID    `json:"reviewer" db:"reviewer"`
	Remarks  string    `json:"remarks" db:"remarks"`
	Added    time.Time `json:"added" db:"added"`
	Edited   time.Time `json:"edited" db:"edited"`
}
