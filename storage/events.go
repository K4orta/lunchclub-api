package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
)

// InsertEvent takes a DB connect and inserts a new EventEntry
func InsertEvent(db *sqlx.DB, event *models.Event) (*models.Event, error) {
	rows, err := db.NamedQuery(db.Rebind(`
		INSERT INTO events (title, slug, start_time, end_time, organizer_id, location_id, rsvps)
		VALUES (:title, :slug, :start_time, :end_time, :organizer_id, :location_id, :rsvps)
		RETURNING id`), event)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&event.ID)
	}

	return event, nil
}
