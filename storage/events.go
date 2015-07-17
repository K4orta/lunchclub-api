package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
)

// InsertEvent takes a DB connect and inserts a new EventEntry
func InsertEvent(db *sqlx.DB, user *models.Event) error {
	_, err := db.NamedExec(db.Rebind(`INSERT INTO events (title, slug, start_time, end_time, organizer_id, location_id, rsvps) VALUES (:title, :slug, :start_time, :end_time, :organizer_id, :location_id, :rsvps)`), user)
	if err != nil {
		return err
	}

	return nil
}
