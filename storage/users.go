// create table users (id SERIAL PRIMARY KEY, fbid text NOT NULL, first_name text NOT NULL, last_name text NOT NULL, permissions text[], added timestamp NOT NULL);
package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
)

func InsertUser(db *sqlx.DB, user *models.User) error {
	_, err := db.NamedExec(db.Rebind(`INSERT INTO users (fbid, first_name, last_name, permissions) VALUES (:fbid, :first_name, :last_name, :permissions)`), user)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(c *sqlx.DB, user *models.User) error {
	return nil
}
