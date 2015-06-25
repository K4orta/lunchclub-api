package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"testing"
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
	`,
	drop: `
	DROP TABLE users;
	`,
}

var pgdb *sqlx.DB

func init() {
	connectionDSN = os.Getenv("LCAPI_TEST_POSTGRES_DSN")
	db, err := CreateConnection()
	if err != nil {
		fmt.Printf("Error connecting to test DB:\n %v\n", err)
	}
	pgdb = db
}

func (s Schema) Postgres() (string, string) {
	return s.create, s.drop
}

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

func runStorageTest(t *testing.T, test func(db *sqlx.DB, t *testing.T)) {
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
