// create table users (id SERIAL PRIMARY KEY, fbid text NOT NULL, first_name text NOT NULL, last_name text NOT NULL, permissions text[], added timestamp NOT NULL);
package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/club"
)

func InsertUser(c *sqlx.DB, user *club.User) {

}

func GetUser(c *sqlx.DB, user *club.User) {

}
