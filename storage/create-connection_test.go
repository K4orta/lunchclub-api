package storage

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	db, err := CreateConnection("lunchclub_test")
	if err != nil {
		t.Error(err)
	}
	db.Close()
}
