package storage

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
	"github.com/k4orta/lunchclub-api/storage/types"
)

func TestInsertUser(t *testing.T) {
	runStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var user = models.User{
			Id:        0,
			FbId:      "abc",
			LastName:  "Wong",
			FirstName: "Erik",
			Roles:     types.StringList{"admin", "mod"},
		}
		InsertUser(db, &user)

		u := models.User{}
		db.Get(&u, db.Rebind(`SELECT * FROM users LIMIT 1`))

		if u.FirstName != "Erik" || u.LastName != "Wong" {
			t.Error("User did not insert name correctly", u)
		}

		if u.Roles[0] != "admin" || u.Roles[1] != "mod" {
			t.Error("User did not insert permissions correctly", u)
		}
	})
}

func TestReadUser(t *testing.T) {
	runStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		InsertUser(db, &models.User{
			Id:        1,
			FbId:      "abc",
			LastName:  "Tio",
			FirstName: "Andrew",
			Roles:     types.StringList{},
		})

		u, _ := GetUserByID(db, 1)

		if u.FirstName != "Andrew" || u.LastName != "Tio" {
			t.Error("Did not fetch user correctly")
		}

	})
}
