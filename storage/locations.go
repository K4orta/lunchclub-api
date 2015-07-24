package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
)

func InsertLocation(db *sqlx.DB, location *models.Location) (*models.Location, error) {
	rows, err := db.NamedQuery(db.Rebind(`
		INSERT INTO locations (name, slug, address, lat_lng)
		VALUES (:name, :slug, :address, :lat_lng)
	`), location)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&location.ID)
	}

	return location, nil
}

// GetLocationBySlug Looks up a location by its Yelp ID
func GetLocationBySlug(db *sqlx.DB, slug string) (*models.Location, error) {
	l := models.Location{}
	err := db.Get(&l, db.Rebind(`SELECT * FROM locations WHERE slug=?`), slug)
	if err != nil {
		log.Printf("Error while getting location by slug: %v", err)
		return nil, err
	}
	return &l, nil
}
