package storage

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
)

// Schema hold the commands to create and destroy the db
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

// Setup the storage module for testing
func SetupDBForTesting() {
	connectionDSN = os.Getenv("LCAPI_TEST_POSTGRES_DSN")
	db, err := CreateConnection()
	if err != nil {
		fmt.Printf("Error connecting to test DB:\n %v\n", err)
	}
	pgdb = db
}

// MultiExec is a helper function for running multiple queries.
func MultiExec(e sqlx.Execer, query string) {
	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := e.Exec(s)
		if err != nil {
			fmt.Println(err, s)
		}
	}
}

func (s Schema) postgres() (string, string) {
	return s.create, s.drop
}

// InjectMockData inserts some sample entries for testing
func InjectMockData() {

}

// RunStorageTest is a utility for running a test with a temporary schema
func RunStorageTest(t *testing.T, test func(db *sqlx.DB, t *testing.T)) {
	runner := func(db *sqlx.DB, t *testing.T, create, drop string) {
		defer func() {
			MultiExec(db, drop)
		}()

		MultiExec(db, create)
		test(db, t)
	}

	create, drop := lcSchema.postgres()
	runner(pgdb, t, create, drop)
}
