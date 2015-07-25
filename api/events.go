package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/k4orta/lunchclub-api/models"
	"github.com/k4orta/lunchclub-api/storage"
	"github.com/k4orta/lunchclub-api/yelp"
)

type eventForm struct {
	YelpURL   string
	StartTime string
}

// CreateEvent handles a POST request to create a new Event and save it in the DB
func CreateEvent(w http.ResponseWriter, req *http.Request) {
	// Read the body
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// Unmarshal the body
	var ef eventForm
	err = json.Unmarshal(b, &ef)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// Connect to the DB
	db, _ := storage.CreateConnection()
	defer db.Close()

	// Get the location from the DB or make one
	loc, err := storage.GetLocationBySlug(db, yelp.ParseURL(ef.YelpURL))
	if err != nil {
		bus, busError := yelp.FetchBusiness(yelp.ParseURL(ef.YelpURL))
		if busError != nil {
			fmt.Println(busError)
		}
		loc, _ = storage.InsertLocation(db, bus)
	}

	//
	ev, _ := storage.InsertEvent(db, &models.Event{
		Title:      "Lunch Club Presents: " + loc.Name,
		LocationID: loc.ID,
	})

	out, _ := json.Marshal(ev)
	fmt.Fprint(w, string(out))
}
