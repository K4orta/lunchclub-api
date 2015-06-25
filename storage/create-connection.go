package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection(dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=ewong password="+os.Getenv("PSPASS")+" dbname="+dbName+" sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}
