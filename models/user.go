package models

import (
	"time"

	"github.com/k4orta/lunchclub-api/storage/types"
)

type UserID int

type User struct {
	ID        UserID           `json:"id" db:"id"`
	FbID      string           `json:"fbId" db:"fbid"`
	FirstName string           `json:"firstName" db:"first_name"`
	LastName  string           `json:"lastName" db:"last_name"`
	Roles     types.StringList `json:"roles" db:"roles"`
	Added     time.Time        `json:"added" db:"added"`
}
