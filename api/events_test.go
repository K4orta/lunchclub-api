package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/lunchclub-api/storage"
)

func TestCreateEvent(t *testing.T) {
	storage.RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		fs := mockAPI(CreateEvent)
		resp, _ := http.Post(fs.URL, "application/json", bytes.NewBuffer([]byte(`
			{
				"yelpURL": "http://www.yelp.com/biz/20-spot-san-francisco"
				"startTime": ""
			}
		`)))
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		if string(b) != "Hello World" {
			t.Errorf("Expected 'Hello World' got %v", string(b))
		}
	})
}
