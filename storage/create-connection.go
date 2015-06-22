package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection() (*sqlx.DB, error) {
	c, err := sqlx.Connect("postgres", "user=ewong password="+os.Getenv("PSPASS")+" dbname=lunchclub sslmode=disable")
	if err != nil {
		return nil, err
	}
	return c, nil
}
