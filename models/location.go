package models

import (
	"github.com/k4orta/lunchclub-api/storage/types"
)

type LocationID int

// Location describes a lunch spot
type Location struct {
	ID      LocationID      `json:"id" db:"id"`
	Name    string          `json:"name" db:"name"`
	Slug    string          `json:"slug" db:"slug"`
	Address string          `json:"address" db:"address"`
	LatLng  types.FloatList `json:"latLng" db:"lat_lng"`
}
