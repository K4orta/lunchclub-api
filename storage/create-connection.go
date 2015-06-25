package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var connectionDSN string = os.Getenv("LCAPI_POSTGRES_DSN")

func CreateConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionDSN)

	if err != nil {
		return nil, err
	}

	return db, nil
}
