package storage

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
	"github.com/k4orta/lunchclub-api/storage/types"
)

func TestInsertEvent(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var event = models.Event{
			ID:          0,
			Title:       "Lunch Club Presents: Pizza Hut",
			Slug:        "pizza-hut-7-16-2015",
			StartTime:   time.Now(),
			EndTime:     time.Now(),
			OrangizerID: 0,
			LocationID:  0,
			RSVPs:       types.IntList{0, 2},
		}
		_, err := InsertEvent(db, &event)
		if err != nil {
			t.Error(err)
		}
		ev := models.Event{}
		db.Get(&ev, db.Rebind(`SELECT * FROM events LIMIT 1`))

		if ev.ID != 1 {
			t.Error("Did not give event an ID after creation")
		}

		if ev.Title != "Lunch Club Presents: Pizza Hut" {
			t.Error("Did not insert Event Title correctly")
		}
	})
}
