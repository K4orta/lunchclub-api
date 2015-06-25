// create table users (id SERIAL PRIMARY KEY, fbid text NOT NULL, first_name text NOT NULL, last_name text NOT NULL, permissions text[], added timestamp NOT NULL);
package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/models"
	"log"
)

func InsertUser(db *sqlx.DB, user *models.User) error {
	_, err := db.NamedExec(db.Rebind(`INSERT INTO users (fbid, first_name, last_name, roles) VALUES (:fbid, :first_name, :last_name, :roles)`), user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(db *sqlx.DB, id int) (*models.User, error) {
	u := models.User{}
	err := db.Get(&u, db.Rebind(`SELECT * FROM users WHERE id=?`), id)
	if err != nil {
		log.Printf("Error while getting user by id: %v", err)
		return nil, err
	}
	return &u, nil
}
