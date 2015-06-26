package models

import (
	"github.com/k4orta/lunchclub-api/storage/types"
)

type LocationID int

type Location struct {
	ID      LocationID      `json:"id" db:"id"`
	LatLng  types.FloatList `json:"latlng" db:"latLng"`
	Name    string          `json:"name" db:"name"`
	Address string          `json:"address" db:"address"`
	URL     string          `json:"url" db:"url"`
}
