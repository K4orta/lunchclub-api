package storage

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Schema struct {
	create string
	drop   string
}

var lcSchema = Schema{
	create: `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		fbid text NOT NULL,
		first_name text NOT NULL,
		last_name text NOT NULL,
		roles text[],
		added timestamp default now()
	);
	CREATE TABLE events (
		id SERIAL PRIMARY KEY,
		title text NOT NULL,
		slug text NOT NULL,
		start_time timestamp default now(),
		end_time timestamp default now(),
		organizer_id int,
		location_id int,
		rsvps int[]
	);
	CREATE TABLE locations (
		id SERIAL PRIMARY KEY,
		name text NOT NULL,
		slug text NOT NULL,
		address text NOT NULL,
		lat_lng double precision[]
	);
	`,
	drop: `
	DROP TABLE users;
	DROP TABLE events;
	DROP TABLE locations;
	`,
}

var pgdb *sqlx.DB

func init() {
	SetDSN(os.Getenv("LCAPI_TEST_POSTGRES_DSN"))
	db, err := CreateConnection()
	if err != nil {
		fmt.Printf("Error connecting to test DB:\n %v\n", err)
	}
	pgdb = db
}

func (s Schema) Postgres() (string, string) {
	return s.create, s.drop
}

// Utility for running a test with a temporary schema
func RunStorageTest(t *testing.T, test func(db *sqlx.DB, t *testing.T)) {
	runner := func(db *sqlx.DB, t *testing.T, create, drop string) {
		defer func() {
			MultiExec(db, drop)
		}()

		MultiExec(db, create)
		test(db, t)
	}

	create, drop := lcSchema.Postgres()
	runner(pgdb, t, create, drop)
}
