package models

import (
	"github.com/k4orta/lunchclub-api/storage/types"
	"time"
)

type RoleList []string

type User struct {
	Id        int              `json:"id" db:"id"`
	FbId      string           `json:"fbId" db:"fbid"`
	FirstName string           `json:"firstName" db:"first_name"`
	LastName  string           `json:"lastName" db:"last_name"`
	Roles     types.StringList `json:"roles" db:"roles"`
	Added     time.Time        `json:"added" db:"added"`
}
