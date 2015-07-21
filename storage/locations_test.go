package storage

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
	"github.com/k4orta/lunchclub-api/storage/types"
)

func TestInsertLocation(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var location = models.Location{
			ID:      0,
			Name:    "Pizza Hut",
			Slug:    "pizza-hut",
			Address: "123 4th Street",
			LatLng:  types.FloatList{0, 2},
		}
		err := InsertLocation(db, &location)
		if err != nil {
			t.Error(err)
		}
		loc := models.Location{}
		db.Get(&loc, db.Rebind(`SELECT * FROM locations LIMIT 1`))

		if loc.Name != "Pizza Hut" {
			t.Error("Did not insert Location Name correctly")
		}
	})
}

func TestGetLocationBySlug(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var location = models.Location{
			ID:      0,
			Name:    "Pizza Hut",
			Slug:    "pizza-hut",
			Address: "123 4th Street",
			LatLng:  types.FloatList{0, 2},
		}
		InsertLocation(db, &location)

		loc, err := GetLocationBySlug(db, "pizza-hut")
		if err != nil {
			t.Error(err)
		}

		if loc.Name != "Pizza Hut" {
			t.Error("Did not select test Location by slug")
		}
	})
}
